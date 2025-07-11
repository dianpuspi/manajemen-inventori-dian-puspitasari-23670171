package controllers

import (
	"inventory-app/config"
	"inventory-app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func Index(c *gin.Context) {
	var items []models.Item
	config.DB.Find(&items)

	var totalBarang int64
	config.DB.Model(&models.Item{}).Select("SUM(stock)").Scan(&totalBarang)

	var totalMasuk int64
	config.DB.Model(&models.BarangMasuk{}).Select("SUM(jumlah)").Scan(&totalMasuk)

	var totalKeluar int64
	config.DB.Model(&models.BarangKeluar{}).Select("SUM(jumlah)").Scan(&totalKeluar)

	// Format waktu sekarang
	now := time.Now().Format("02 Jan 2006 15:04")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"items":       items,
		"totalBarang": totalBarang,
		"totalMasuk":  totalMasuk,
		"totalKeluar": totalKeluar,
		"lastUpdated": now, // dikirim ke template
	})
}

func ShowForm(c *gin.Context) {
	var items []models.Item
	config.DB.Find(&items)

	var totalBarang int64
	config.DB.Model(&models.Item{}).Select("SUM(stock)").Scan(&totalBarang)

	c.HTML(http.StatusOK, "form.html", gin.H{
		"items":       items,
		"totalBarang": totalBarang,
	})
}

func Create(c *gin.Context) {
	name := c.PostForm("name")
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)

	item := models.Item{Name: name, Stock: stock, Price: price}
	config.DB.Create(&item)

	// Ambil data ulang untuk ditampilkan
	var items []models.Item
	var totalBarang int64

	config.DB.Find(&items)
	config.DB.Model(&models.Item{}).Select("SUM(stock)").Scan(&totalBarang)

	c.HTML(http.StatusOK, "form.html", gin.H{
		"items":       items,
		"totalBarang": totalBarang,
	})
}

func Edit(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	config.DB.First(&item, id)
	c.HTML(http.StatusOK, "form.html", gin.H{"item": item})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	config.DB.First(&item, id)
	item.Name = c.PostForm("name")
	item.Stock, _ = strconv.Atoi(c.PostForm("stock"))
	item.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	config.DB.Save(&item)
	c.Redirect(http.StatusFound, "/new")
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Item{}, id)
	c.Redirect(http.StatusFound, "/new")
}

func ExportExcel(c *gin.Context) {
	var items []models.Item
	config.DB.Find(&items)

	f := excelize.NewFile()
	f.SetSheetRow("Sheet1", "A1", &[]string{"ID", "Nama", "Stock", "Harga", "Tanggal"})
	for i, item := range items {
		row := []interface{}{item.ID, item.Name, item.Stock, item.Price, item.CreatedAt}
		f.SetSheetRow("Sheet1", "A"+strconv.Itoa(i+2), &row)
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=inventory.xlsx")
	f.Write(c.Writer)
}
