package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"time"

	"tegar/pcc/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DataLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func buatToken(id uint, nama string) (string, error) {
	var err error

	dataJwt := jwt.MapClaims{}
	dataJwt["authorized"] = true
	dataJwt["id"] = id
	dataJwt["nama"] = nama
	dataJwt["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, dataJwt)
	token, err := at.SignedString([]byte(os.Getenv("SECRET_KEY_JWT")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var vDataLogin DataLogin
	if err := c.ShouldBindJSON(&vDataLogin); err != nil {
		c.JSON(500, gin.H{
			"kode_error": "JSON-ERROR",
			"pesan":      "Format JSON Salah",
		})
		return
	}

	h := sha1.New()
	h.Write([]byte(vDataLogin.Password))
	enkripsiPass := hex.EncodeToString(h.Sum(nil))

	var pengguna models.User
	db.Where("username = ? and password = ?",
		vDataLogin.Username,
		enkripsiPass,
	).First(&pengguna)

	if pengguna.Id == 0 {
		c.JSON(401, gin.H{
			"kode_error": "401",
			"pesan":      "username atau password salah",
		})
		return
	}

	token, err := buatToken(pengguna.Id, pengguna.Nama)

	if err != nil {
		c.JSON(500, gin.H{
			"kode_error": "TOKEN-ERROR",
			"pesan":      "Gagal membuat token JWT",
			"error":      err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"kode_error": "0",
		"pesan":      "berhasil login",
		"token":      token,
	})
}
