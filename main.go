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
	r.GET("/mahasiswa", controller.ReadData)
	r.POST("/insert", controller.CreateData)
	r.Run()
}
