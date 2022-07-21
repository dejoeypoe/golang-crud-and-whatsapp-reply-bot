package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"tegar/pcc/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Upload(c *gin.Context) {
	fileName := c.PostForm("fileName")
	file, _ := c.FormFile("file")
	mimeType := file.Header.Get("Content-Type")
	fileOpen, _ := file.Open()
	defer fileOpen.Close()
	fileData, _ := ioutil.ReadAll(fileOpen)

	// Encode as base64.
	data := base64.StdEncoding.EncodeToString(fileData)
	postBody, _ := json.Marshal(map[string]string{
		"fileName": fileName,
		"mimeType": mimeType,
		"data":     data,
	})
	requestBody := bytes.NewBuffer(postBody)

	// post data
	res, err := http.Post(
		"https://script.google.com/macros/s/AKfycbzPF5Uwa3D6ofoYLAbl7TiAJmC3FKaUoiQ_risVWQBLOD0niFl6pWoBBZKjLhe7dXgX4w/exec",
		"application/json; charset=UTF-8",
		requestBody,
	)

	// apakah ada error
	if err != nil {
		c.JSON(500, gin.H{
			"kode_error": "ERR-DRIVE",
			"pesan":      "Gagal Upload",
		})
		return
	}
	//baca response data
	hasilBody, _ := ioutil.ReadAll(res.Body)
	hasilString := string(hasilBody)

	//konversi string json to object
	var hasilJson map[string]interface{}
	json.Unmarshal([]byte(hasilString), &hasilJson)
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil Upload",
		"data":       hasilJson,
	})

	// close response body
	res.Body.Close()
}

//tampil data dari tabel
func Tampil(c *gin.Context) {
	//ambil koneksi ke variabel db
	db := c.MustGet("db").(*gorm.DB)
	//membuat variabel dokumen berupa array dari model Dokumen
	var dokumen []models.Dokumen
	//ambil semua data dari tabel
	db.Find(&dokumen)
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil Tampil",
		"data":       dokumen,
	})
}

//direct download dari google drive
func Unduh(c *gin.Context) {
	id := c.Param("id")
	//get data
	res, err := http.Get("https://script.google.com/macros/s/AKfycbzPF5Uwa3D6ofoYLAbl7TiAJmC3FKaUoiQ_risVWQBLOD0niFl6pWoBBZKjLhe7dXgX4w/exec?id=" + id)
	//apakah ada error
	if err != nil {
		c.JSON(500, gin.H{
			"kode_error": "ERR-DRIVE",
			"pesan":      "Gagal Unduh",
		})
		return
	}
	//baca response data
	hasilBody, _ := ioutil.ReadAll(res.Body)
	hasilString := string(hasilBody)
	//konversi string json to object
	var hasilJson map[string]interface{}
	json.Unmarshal([]byte(hasilString), &hasilJson)
	//ambil data file dan mimeType
	fileBase64 := hasilJson["file"].(string)
	mimeType := hasilJson["mimeType"].(string)
	// konversi base64 ke file
	file, _ := base64.StdEncoding.DecodeString(fileBase64)
	//tulis dengan header sesuai dengan mimeType
	c.Writer.Header().Set("Content-Type", mimeType)
	c.Writer.Write(file)
}
