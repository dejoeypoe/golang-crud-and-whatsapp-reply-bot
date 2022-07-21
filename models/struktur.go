package models

import (
	"time"
)

type Dokumen struct {
	Id          uint      `json:"id" gorm:"prrimary_key"`
	NamaDokumen string    `json:"nama_dokumen" gorm:"type:varchar(150)"`
	FileId      string    `json:"file_id" gorm: "type:varchar(255)"`
	FileUrl     string    `json:"file_url" gorm: "type:varchar(255)"`
	WaktuUpload time.Time `json:"waktu_upload" gorm: "not null;default:CURRENT_TIMESTAMP"`
}

type User struct {
	Id       uint   `json: "id"  gorm:"primary_key"`
	Nama     string `json:"nama" gorm:"type:varchar(150)"`
	Username string `json:"username" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
}

type Pegawai struct {
	Nik    string `json:"nik" gorm:"type:varchar(50); primary_key"`
	Nama   string `json:"nama" gorm:"type:varchar(150)"`
	Jk     string `json:"jk" gorm:"type:varchar(100)"`
	NoTelp string `json:"no_telp" gorm:"type:varchar(100)"`
	Alamat string `json:"alamat" gorm:"type:varchar(255)"`
}

type Jabatan struct {
	Id   uint   `json:"id" gorm:"type:int(8); primary_key"`
	Nama string `json:"nama" gorm:"type:varchar(100)"`
}

type RiwayatJabatan struct {
	Id        uint   `json: "id" gorm:""type:int(8); primary_key"`
	Nik       string `json: "nik" gorm:"type:varchar(50)"`
	JabatanId uint   `json: "jabatan_id" gorm:""type:int(8)"`
}

type Cuti struct {
	Id           uint   `json:"id" gorm:"primary_key"`
	Nik          string `json:"nik" gorm:"type:varchar(50)"`
	TanggalAwal  string `json:"tanggal_awal" gorm:"type:date"`
	TanggalAkhir string `json:"tanggal_akhir" gorm:"type:date"`
	PotonganGaji int    `json:"potongan_gaji" gorm:"type:integer(11)"`
}

type Chat struct {
	Kode    string `json:"kode" gorm:"type:varchar(20);primary_key"`
	Balasan string `json:"balasan" gorm:"type:text"`
}
