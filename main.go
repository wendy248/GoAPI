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

	//Create
	r.POST("/mahasiswa", controller.CreateData)

	//Update
	r.PUT("/test/:nim", controller.UpdateData)

	//Delete
	r.DELETE("/mahasiswa", controller.DeleteData)
	r.Run()
}
