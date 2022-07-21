package controllers

import (
	"tegar/pcc/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//fungsi tampil data
func RiwayatJabatanTampil(c *gin.Context) {

	//ambil koneksi db yg dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat objek array dr model RiwayatJabatan
	var arrayRiwayatJabatan []models.RiwayatJabatan

	//ambil semua records dr tabel RiwayatJabatan & masukkan ke arrayRiwayatJabatan
	hasil := db.Find(&arrayRiwayatJabatan)

	//jika ada kesalahan tampil
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal tampil data",
			"kesalahan":  hasil.Error,
		})
		return
	}

	//tampilkan respon json
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil Tampil",
		"data":       arrayRiwayatJabatan,
		"count":      hasil.RowsAffected,
	})
}

//fungsi tambah data
func RiwayatJabatanTambah(c *gin.Context) {

	//ambil koneksi db yg dibawa dr main.gorm
	db := c.MustGet("db").(*gorm.DB)

	//membuat objek sr model mhs
	var inputRiwayatJabatan models.RiwayatJabatan

	//validasi input data yg dikirim dan disimpan di var inputRiwayatJabatan
	errorInput := c.ShouldBindJSON(&inputRiwayatJabatan)

	//jika ada kesalahan input tampilkan pesan
	if errorInput != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Terdapat kesalalhan pada data JSON yang dikirim",
			"kesalahan":  errorInput,
		})
		return
	}

	//buat record baru dr inputRiwayatJabatan
	hasil := db.Create(&inputRiwayatJabatan)

	//jika ada error saat buat data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      " Gagal menambahkan data",
			"data":       inputRiwayatJabatan,
			"kesalahan":  hasil.Error,
		})
		return
	}

	//tampilkan repon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil menambahkan data",
		"data":       hasil.Value,
	})

}

//fungsi ubah data
func RiwayatJabatanUbah(c *gin.Context) {

	//ambil koneksi db yg dibawa dr main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat objek dr model mhs
	var inputRiwayatJabatan models.RiwayatJabatan

	//validasi input data yg dikirim dan simpan di var inputRiwayatJabatan
	errorInput := c.ShouldBindJSON(&inputRiwayatJabatan)

	//jika ada kesalahan input tampilkan pesan
	if errorInput != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Terdapat kesalahan pada data JSON yg dikirim",
			"kesalahan":  errorInput,
		})
		return
	}

	//ubah record
	hasil := db.Model(&models.RiwayatJabatan{Id: inputRiwayatJabatan.Id}).Updates(inputRiwayatJabatan)

	//jika ada error saat ubah data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal ubah data",
			"data":       inputRiwayatJabatan,
			"kesalahan":  hasil.Error,
		})
		return
	}

	//tampiljan respon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil ubah",
		"data":       hasil.Value,
		"count":      hasil.RowsAffected,
	})
}

//fungsi hapus data
func RiwayatJabatanHapus(c *gin.Context) {

	//ambil koneksi db yg dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membaca parameter id dari url
	id := c.Param("id")

	//hapus record
	hasil := db.Delete(&models.RiwayatJabatan{}, id)

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

	//tampil respon JSON
	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "Berhasil hapus",
		"data":       id,
		"count":      hasil.RowsAffected,
	})
}
