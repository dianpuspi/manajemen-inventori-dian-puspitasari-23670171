package main

import (
	"encoding/json"
	"html/template"
	"inventory-app/config"
	"inventory-app/controllers"
	"inventory-app/models"

	"github.com/gin-gonic/gin"
)

// Fungsi helper untuk konversi data ke JSON agar bisa digunakan di template HTML
func toJson(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}

func main() {
	// Koneksi ke database
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Item{}, &models.BarangMasuk{}, &models.BarangKeluar{})

	// Inisialisasi Gin
	r := gin.Default()

	// Gunakan template dengan fungsi custom toJson
	tmpl := template.New("").Funcs(template.FuncMap{
		"toJson": toJson,
	})
	tmpl = template.Must(tmpl.ParseGlob("templates/*"))
	r.SetHTMLTemplate(tmpl)

	// Routing
	r.GET("/", controllers.Index)
	r.GET("/new", controllers.ShowForm)
	r.POST("/create", controllers.Create)
	r.GET("/edit/:id", controllers.Edit)
	r.POST("/update/:id", controllers.Update)
	r.GET("/delete/:id", controllers.Delete)
	r.GET("/export", controllers.ExportExcel)

	r.GET("/barang-masuk", controllers.FormBarangMasuk)
	r.POST("/barang-masuk", controllers.ProsesBarangMasuk)
	r.GET("/barang-keluar", controllers.FormBarangKeluar)
	r.POST("/barang-keluar", controllers.ProsesBarangKeluar)

	// Jalankan server di port 8080
	r.Run(":8080")
}
