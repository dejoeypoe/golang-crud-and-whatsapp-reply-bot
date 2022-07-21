package controllers

import (
	"tegar/pcc/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//fugsi tampil data
func CutiTampil(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat object array dari model pegawai
	var arrayCuti []models.Cuti

	//ambil semua records dari tabel pegawai dan masukkan ke arrayPegawai
	hasil := db.Find(&arrayCuti)

	//jika ada kesalahan tampil
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal tampil data",
			"kesalahan":  hasil.Error,
		})
		return
	}

	//tampilkan respon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil tampil",
		"data":       arrayCuti,
		"count":      hasil.RowsAffected,
	})
}

func CutiTampilBerdasarkan(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	nik := c.Param("nik")
	tanggal_awal := c.Param("tanggal_awal")
	//membuat object array dari model pegawai
	var arrayCuti []models.Cuti

	//ambil semua records dari tabel pegawai dan masukkan ke arrayPegawai
	hasil := db.Where("nik = ? AND tanggal_awal =?", nik, tanggal_awal).First(&arrayCuti)

	//jika ada kesalahan tampil
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal tampil data",
			"kesalahan":  hasil.Error,
		})
		return

	}

	//tampilkan respon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil tampil",
		"data":       arrayCuti,
		"count":      hasil.RowsAffected,
	})

}

//fungsi tambah data
func CutiTambah(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat object array dari model pegawai
	var inputCuti models.Cuti

	//validasi input data yang dikirim dan simpan
	errorInput := c.ShouldBindJSON(&inputCuti)

	//jika ada kesalahan input tampilkan pesan
	if errorInput != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Terdapat kesalahan pada data JSON yang dikirim",
			"kesalahan":  errorInput,
		})
		return
	}

	//buat record baru dari inputPegawai
	hasil := db.Create(&inputCuti)

	//jika ada error saat buat data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal tambah data",
			"data":       inputCuti,
			"kesalahan":  hasil.Error,
		})
		return
	}

	//tampilkan respon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil Tambah",
		"data":       hasil.Value,
	})
}

//fungsi ubah data
func CutiUbah(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat object array dari model pegawai
	var inputCuti models.Cuti

	//validasi input data yang dikirim dan simpan
	errorInput := c.ShouldBindJSON(&inputCuti)

	//jika ada kesalahan input tampilkan pesan
	if errorInput != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Terdapat kesalahan pada data JSON yang dikirim",
			"kesalahan":  errorInput,
		})
		return
	}

	//ubah record
	hasil := db.Model(&models.Cuti{Id: inputCuti.Id}).Updates(inputCuti)

	//jika ada error saat buat data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal ubah data",
			"data":       inputCuti,
			"kesalahan":  hasil.Error,
		})
		return
	}

	//tampilkan respon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil ubah",
		"data":       hasil.Value,
		"count":      hasil.RowsAffected,
	})
}

//fungsi hapus data
func CutiHapus(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membaca parameter nik dari url
	id := c.Param("id")

	//hapus record
	hasil := db.Delete(&models.Cuti{}, id)

	//jika ada error saat hapus data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal hapus data",
			"data":       id,
			"kesalahan":  hasil.Error,
		})
		return
	}

	//tampilkan respon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil hapus",
		"data":       id,
		"count":      hasil.RowsAffected,
	})
}
