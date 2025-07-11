package controllers

import (
	"inventory-app/config"
	"inventory-app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func FormBarangMasuk(c *gin.Context) {
	var items []models.Item
	var barangMasuk []models.BarangMasuk
	var totalBarang int64
	var totalMasuk int64

	config.DB.Find(&items)
	config.DB.Preload("Item").Find(&barangMasuk)
	config.DB.Model(&models.Item{}).Select("SUM(stock)").Scan(&totalBarang)
	config.DB.Model(&models.BarangMasuk{}).Select("SUM(jumlah)").Scan(&totalMasuk)

	c.HTML(http.StatusOK, "barang_masuk.html", gin.H{
		"items":       items,
		"barangMasuk": barangMasuk,
		"totalBarang": totalBarang,
		"totalMasuk":  totalMasuk,
	})
}

func ProsesBarangMasuk(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.PostForm("item_id"))
	jumlah, _ := strconv.Atoi(c.PostForm("jumlah"))
	masuk := models.BarangMasuk{
		ItemID:  uint(itemID),
		Jumlah:  jumlah,
		Tanggal: time.Now(), // Set tanggal saat ini
	}
	config.DB.Create(&masuk)

	var item models.Item
	config.DB.First(&item, itemID)
	item.Stock += jumlah
	config.DB.Save(&item)

	FormBarangMasuk(c)
}

func FormBarangKeluar(c *gin.Context) {
	var items []models.Item
	var keluarList []models.BarangKeluar

	config.DB.Find(&items)
	config.DB.Preload("Item").Find(&keluarList)

	var totalBarang int64
	var totalKeluar int64
	config.DB.Model(&models.Item{}).Select("SUM(stock)").Scan(&totalBarang)
	config.DB.Model(&models.BarangKeluar{}).Select("SUM(jumlah)").Scan(&totalKeluar)

	c.HTML(http.StatusOK, "barang_keluar.html", gin.H{
		"items":        items,
		"barangKeluar": keluarList,
		"totalBarang":  totalBarang,
		"totalKeluar":  totalKeluar,
	})
}

func ProsesBarangKeluar(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.PostForm("item_id"))
	jumlah, _ := strconv.Atoi(c.PostForm("jumlah"))

	today := time.Now().In(time.FixedZone("WIB", 7*3600)).Truncate(24 * time.Hour)

	keluar := models.BarangKeluar{
		ItemID:  uint(itemID),
		Jumlah:  jumlah,
		Tanggal: today,
	}

	config.DB.Create(&keluar)

	var item models.Item
	config.DB.First(&item, itemID)
	if item.Stock >= jumlah {
		item.Stock -= jumlah
		config.DB.Save(&item)
	}

	var items []models.Item
	var keluarList []models.BarangKeluar
	config.DB.Find(&items)
	config.DB.Preload("Item").Find(&keluarList)

	var totalBarang int64
	var totalKeluar int64
	config.DB.Model(&models.Item{}).Select("SUM(stock)").Scan(&totalBarang)
	config.DB.Model(&models.BarangKeluar{}).Select("SUM(jumlah)").Scan(&totalKeluar)

	c.HTML(http.StatusOK, "barang_keluar.html", gin.H{
		"items":        items,
		"barangKeluar": keluarList,
		"totalBarang":  totalBarang,
		"totalKeluar":  totalKeluar,
	})
}
