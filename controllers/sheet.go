package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//tampil data dari tabel
func TampilSheet(c *gin.Context) {
	res, err := http.Get(
		"https://script.google.com/macros/s/AKfycbwH85vtjUQUtRQ-R2PzjaFKOrKHSNZUqMAeHbOIlfFdi60MaYnWOCkWJYENZvzqMkybfQ/exec",
	)
	// apakah ada error
	if err != nil {
		c.JSON(500, gin.H{
			"kode_error": "ERR-SHEET",
			"pesan":      "Gagal Baca Sheet",
		})
		return
	}
	// baca response data
	hasilBody, _ := ioutil.ReadAll(res.Body)
	hasilString := string(hasilBody)
	//konversi string json to object
	var hasilJson map[string]interface{}
	json.Unmarshal([]byte(hasilString), &hasilJson)
	//tampilan json
	c.JSON(200, hasilJson)
	// close response body
	res.Body.Close()
}
