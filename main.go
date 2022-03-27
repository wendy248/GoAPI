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

	v1 := r.Group("/api/v1")
	//Read
	v1.GET("/mahasiswa", controller.ReadData)
	v1.GET("/matakuliah", controller.ReadMatkul)

	//Create
	v1.POST("/mahasiswa", controller.CreateData)
	v1.POST("/matakuliah", controller.CreateMatkul)

	//Update
	v1.PUT("/mahasiswa/:nim", controller.UpdateData)
	v1.PUT("/matakuliah/:kode", controller.UpdateMatkul)

	//Delete
	v1.DELETE("/mahasiswa", controller.DeleteData)
	v1.DELETE("/matakuliah", controller.DeleteMatkul)
	r.Run()
}
