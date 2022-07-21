package controllers

import (
	"tegar/pcc/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//fungsi tampil data
func JabatanTampil(c *gin.Context) {

	//ambil koneksi db yg dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat objek array dr model Jabatan
	var arrayJabatan []models.Jabatan

	//ambil semua records dr tabel Jabatan & masukkan ke arrayJabatan
	hasil := db.Find(&arrayJabatan)

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
		"data":       arrayJabatan,
		"count":      hasil.RowsAffected,
	})
}

//fungsi tambah data
func JabatanTambah(c *gin.Context) {

	//ambil koneksi db yg dibawa dr main.gorm
	db := c.MustGet("db").(*gorm.DB)

	//membuat objek sr model mhs
	var inputJabatan models.Jabatan

	//validasi input data yg dikirim dan disimpan di var inputJabatan
	errorInput := c.ShouldBindJSON(&inputJabatan)

	//jika ada kesalahan input tampilkan pesan
	if errorInput != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Terdapat kesalalhan pada data JSON yang dikirim",
			"kesalahan":  errorInput,
		})
		return
	}

	//buat record baru dr inputJabatan
	hasil := db.Create(&inputJabatan)

	//jika ada error saat buat data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      " Gagal menambahkan data",
			"data":       inputJabatan,
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
func JabatanUbah(c *gin.Context) {

	//ambil koneksi db yg dibawa dr main.go
	db := c.MustGet("db").(*gorm.DB)

	//membuat objek dr model mhs
	var inputJabatan models.Jabatan

	//validasi input data yg dikirim dan simpan di var inputJabatan
	errorInput := c.ShouldBindJSON(&inputJabatan)

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
	hasil := db.Model(&models.Jabatan{Id: inputJabatan.Id}).Updates(inputJabatan)

	//jika ada error saat ubah data
	if hasil.Error != nil {
		c.JSON(500, gin.H{
			"kode_error": "500",
			"pesan":      "Gagal ubah data",
			"data":       inputJabatan,
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
func JabatanHapus(c *gin.Context) {

	//ambil koneksi db yg dibawa dari main.go
	db := c.MustGet("db").(*gorm.DB)

	//membaca parameter id dari url
	id := c.Param("id")

	//hapus record
	hasil := db.Delete(&models.Jabatan{}, id)

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
