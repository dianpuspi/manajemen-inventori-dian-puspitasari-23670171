package models

import "time"

type BarangMasuk struct {
	ID      uint
	ItemID  uint
	Item    Item
	Jumlah  int
	Tanggal time.Time
}

type BarangKeluar struct {
	ID      uint
	ItemID  uint
	Item    Item
	Jumlah  int
	Tanggal time.Time `gorm:"type:date"`
}
