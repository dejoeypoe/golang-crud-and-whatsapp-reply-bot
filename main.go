package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"tegar/pcc/controllers"
	"tegar/pcc/middleware"
	"tegar/pcc/models"
	"tegar/pcc/wa"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Mengaktifkan mode release/production
	gin.SetMode(gin.ReleaseMode)

	// Memvaca file .env
	godotenv.Load()

	// Setting Zona waktu
	os.Setenv("TZ", "Asia/Jakarta")

	// Baca waktu saat ini sesuai dengan komputer
	waktuSaiki := time.Now()

	// Membuat file log request
	nama_file_request := fmt.Sprint("log/request_", waktuSaiki.Format("2006-01-02 15-04-05"), ".log")
	file_request, _ := os.Create(nama_file_request)
	gin.DefaultWriter = io.MultiWriter(file_request)

	// Membuat file log error
	nama_file_error := fmt.Sprint("log/error_", waktuSaiki.Format("2006-01-02 15-04-05"), ".log")
	file_error, _ := os.Create(nama_file_error)
	gin.DefaultErrorWriter = io.MultiWriter(file_error)

	// Koneksi
	db := models.Koneksi()
	nama_db_error := fmt.Sprint("log/db_error_", waktuSaiki.Format("2006-01-02 15-04-05"), ".log")
	db_error, _ := os.Create(nama_db_error)
	db.SetLogger(log.New(db_error, "", log.LstdFlags|log.Lshortfile))

	// Membuat Object Gin Gonic Dengan nama: r
	r := gin.New()

	// Mengaktifkan Log
	r.Use(gin.Logger())

	// Mengaktifkan proses recovery jika ada eror
	r.Use(gin.Recovery())

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "x-csrf-token", "_token"}
	r.Use(cors.New(config))

	// Jika tidak ditemukan route saat akses
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"kode_error": "404",
			"pesan":      "Route Tidak Ada",
		})
	})

	// Mendaftarkan koneksi ke gin gonic
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Route utama
	r.Use(middleware.AksesMiddleware())
	r.GET("/", func(e *gin.Context) {
		e.JSON(200, gin.H{
			"kode_error": "0",
			"Pesan":      "Selamat datang di Universitas Duta Bangsa",
		})
	})

	// Method Post Login
	r.POST("/login", controllers.Login)

	// Route API
	api := r.Group("/api")
	api.Use(middleware.JwtMiddleware())
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"kode_error": "0",
			"pesan":      "Selamat Datang di API FIKOM UDB",
		})
	})

	//route pegawai
	api.GET("/pegawai", controllers.PegawaiTampil)
	api.POST("/pegawai", controllers.PegawaiTambah)
	api.PUT("/pegawai", controllers.PegawaiUbah)
	api.DELETE("/pegawai/:nik", controllers.PegawaiHapus)

	//route jabatan
	api.GET("/jabatan", controllers.JabatanTampil)
	api.POST("/jabatan", controllers.JabatanTambah)
	api.PUT("/jabatan", controllers.JabatanUbah)
	api.DELETE("/jabatan/:id", controllers.JabatanHapus)

	//route riwayat jabatan
	api.GET("/riwayat_jabatan", controllers.RiwayatJabatanTampil)
	api.POST("/riwayat_jabatan", controllers.RiwayatJabatanTambah)
	api.PUT("/riwayat_jabatan", controllers.RiwayatJabatanUbah)
	api.DELETE("/riwayat_jabatan/:id", controllers.RiwayatJabatanHapus)

	//route Cuti
	api.GET("/cuti", controllers.CutiTampil)
	api.GET("/cuti/:nik/:tanggal_awal", controllers.CutiTampilBerdasarkan)
	api.POST("/cuti", controllers.CutiTambah)
	api.PUT("/cuti", controllers.CutiUbah)
	api.DELETE("/cuti/:id", controllers.CutiHapus)

	//route drive
	api.POST("/drive", controllers.Upload)
	api.GET("/drive", controllers.Tampil)
	api.GET("/drive/:id", controllers.Unduh)

	//route sheet
	r.GET("/sheet", controllers.TampilSheet)

	//route wa
	go func() {
		wa.InitWa()
	}()

	// Membaca variabel PORT di .env
	port := os.Getenv("PORT")

	// Tampilkan pesan di cmd
	fmt.Println("Server Berjalan Di Port : ", port)

	// Jalankan server di port yang ditentukan di .env
	r.Run(":" + port)

}
