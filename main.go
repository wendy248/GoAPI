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

	//Create
	v1.POST("/mahasiswa", controller.CreateData)

	//Update
	v1.PUT("/test/:nim", controller.UpdateData)

	//Delete
	// v1.DELETE("/mahasiswa")
	r.Run()
}
