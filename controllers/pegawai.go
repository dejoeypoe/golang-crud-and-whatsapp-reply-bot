package controllers

import (
	"tegar/pcc/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//fugsi tampil data
func PegawaiTampil(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat object array dari model pegawai
	var arrayPegawai []models.Pegawai

	//ambil semua records dari tabel pegawai dan masukkan ke arrayPegawai
	hasil := db.Find(&arrayPegawai)

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
		"data":       arrayPegawai,
		"count":      hasil.RowsAffected,
	})
}

//fungsi tambah data
func PegawaiTambah(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat object array dari model pegawai
	var inputPegawai models.Pegawai

	//validasi input data yang dikirim dan simpan
	errorInput := c.ShouldBindJSON(&inputPegawai)

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
	hasil := db.Create(&inputPegawai)

	//jika ada error saat buat data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal tambah data",
			"data":       inputPegawai,
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
func PegawaiUbah(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat object array dari model pegawai
	var inputPegawai models.Pegawai

	//validasi input data yang dikirim dan simpan
	errorInput := c.ShouldBindJSON(&inputPegawai)

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
	hasil := db.Model(&models.Pegawai{Nik: inputPegawai.Nik}).Updates(inputPegawai)

	//jika ada error saat buat data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal ubah data",
			"data":       inputPegawai,
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
func PegawaiHapus(c *gin.Context) {
	//ambil koneksi db yang dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membaca parameter nik dari url
	nik := c.Param("nik")

	//hapus record
	hasil := db.Delete(&models.Pegawai{}, "nik=?", nik)

	//jika ada error saat hapus data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal hapus data",
			"data":       nik,
			"kesalahan":  hasil.Error,
		})
		return
	}

	//tampilkan respon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil hapus",
		"data":       nik,
		"count":      hasil.RowsAffected,
	})
}
