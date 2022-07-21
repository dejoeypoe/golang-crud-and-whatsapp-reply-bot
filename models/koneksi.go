package models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Koneksi() *gorm.DB {
	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", dbuser+":"+dbpass+"@("+dbhost+")/"+dbname+
		"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	// Auto Migrate
	db.AutoMigrate(&Dokumen{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Pegawai{})

	return db
}
