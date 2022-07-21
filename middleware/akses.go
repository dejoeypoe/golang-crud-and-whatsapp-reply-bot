package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Fungsi untuk menampilkan pesan jika ada kesalahan
func responKesalahan(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"kode_error": "401",
		"pesan":      message,
	})
}

// Fungsi core dari middleware
func AksesMiddleware() gin.HandlerFunc {
	// Baca variable FIKOM_TOKEN di .env
	requiredToken := os.Getenv("FIKOM_TOKEN")

	// Cek apakah FIKOM_TOKEN ada nilainya di .env
	if requiredToken == "" {
		log.Fatal("Silahkan Set FIKOM_TOKEN di file .env")
	}
	return func(c *gin.Context) {

		// Baca Header _token
		token := c.Request.Header.Get("_token")

		// jika _token kosong
		if token == "" {
			responKesalahan(c, 401, "API Token tidak ada")
			return
		}

		// Jika _token tidak sama dengan FIKOMDB_TOKEN
		if token != requiredToken {
			responKesalahan(c, 401, "API Token Tidak Valid")
			return
		}

		// Jika tidak ada error dan _token sama dengan FIKOM_TOKEN maka Request dilanjutkan
		c.Next()
	}

}
