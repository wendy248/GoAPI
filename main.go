package main

import (
	"GoAPI/controller"
	"GoAPI/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// v1 := r.Group("/api/v1")
	//Read
	r.GET("/mahasiswa", controller.ReadData)
	r.GET("/matakuliah", controller.ReadMatkul)

	//Create
	r.POST("/mahasiswa", controller.CreateData)
	r.POST("/matakuliah", controller.CreateMatkul)

	//Update
	r.PUT("/mahasiswa/:nim", controller.UpdateData)
	r.PUT("/matakuliah/:kode", controller.UpdateMatkul)

	//Delete
	r.DELETE("/mahasiswa", controller.DeleteData)
	r.Run()
}
