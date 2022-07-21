package wa

import (
	"context"
	"fmt"

	"os"
	"time"

	"io/ioutil"

	"tegar/pcc/models"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

var client *whatsmeow.Client

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		if !v.Info.IsFromMe {
			if v.Message.GetConversation() != "" {
				fmt.Println("PESAN DITERIMA!", v.Message.GetConversation())
				balasChat(v)
				/*
					 client.SendMessage(v.Info.Sender, "", &waProto.Message{
					Conversation: proto.String("Pesan ini pesan otomatis. " +
					 v.Message.GetConversation()),
					 })
				*/
			}
		}
	case *events.LoggedOut:
		TutupWa()
		time.Sleep(10 * time.Second)
		InitWa()
	}
}

func InitWa() {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				fmt.Println("QR code:", evt.Code)
				tulisQr([]byte(evt.Code))
			} else if evt.Event == "timeout" {
				fmt.Println("Login event:", evt.Event)
				tulisQr([]byte(""))
				TutupWa()
				time.Sleep(10 * time.Second)
				InitWa()
			} else {
				fmt.Println("Login event:", evt.Event)
				tulisQr([]byte(""))
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

}

//fungsi untuk memutus koneksi dengan server wa
func TutupWa() {
	client.Disconnect()
}

//fungsi untuk menulis kode qr di file qr.wa
func tulisQr(s []byte) {
	f, _ := os.OpenFile("qr.wa", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	f.Write(s)
	f.Close()
}

//fungsi untuk membalas wa dari tabel chat
func balasChat(v *events.Message) {
	//panggil koneksi
	db := models.Koneksi()
	//buat variable chat dari struktur Chat
	var chat models.Chat
	//ambil pesan wa dari pengirim dan simpan di variable kode
	kode := v.Message.GetConversation()
	//ambil data dari tabel chats berdasarkan kode
	db.Where("kode = ?", kode).First(&chat)
	//jika data chat dari tabel tidak kosong (ada datanya)
	if chat.Kode != "" {
		//menandai sudah dibaca
		ids := []string{v.Info.ID}
		client.MarkRead(ids, time.Now(), v.Info.Chat, v.Info.Sender)
		//memantau aktifitas pengirim
		client.SubscribePresence(v.Info.Sender)
		//kirim status available
		client.SendPresence(types.PresenceAvailable)
		//diam dulu 1 detik biar lebih manuasiawi, heheheh
		time.Sleep(1 * time.Second)
		//kirim status sedang mengetik
		client.SendChatPresence(v.Info.Sender,
			types.ChatPresenceComposing, types.ChatPresenceMediaText)
		//diam dulu 2 detik biar lebih manuasiawi, heheheh
		time.Sleep(2 * time.Second)
		//kirim status berhenti mengetik
		client.SendChatPresence(v.Info.Sender,
			types.ChatPresencePaused, types.ChatPresenceMediaText)
		//kirim balasan pesan wa
		client.SendMessage(context.Background(), v.Info.Sender, "", &waProto.Message{Conversation: proto.String(chat.Balasan)})
	}
}

func BacaQr(c *gin.Context) {
	//buka file qr.wa
	f, _ := os.Open("qr.wa")
	//baca isi f
	qr, _ := ioutil.ReadAll(f)
	var pesan string
	//jika Panjang qr > 0 yang artinya ada isinya
	if len(qr) > 0 {
		pesan = "Ada QR Code"
	} else {
		pesan = "Mungkin anda belum terhubung atau mungkin sudah berhasil login"
	}
	//response json
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      pesan,
		"qr":         qr,
	})
}
