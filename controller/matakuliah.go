package controller

import (
	"GoAPI/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type MataKuliahInput struct {
	KodeMatkul    string `json:"kode" binding:"required"`
	NamaMatkul    string `json:"nama" binding:"required"`
	JumlahSKS     int16  `json:"jumlah" binding:"required"`
	DosenPengampu string `json:"dosen" binding:"required"`
}

//ReadData
func ReadMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var matakuliah []models.MataKuliah
	db.Find(&matakuliah)

	c.JSON(http.StatusOK, gin.H{
		"data": matakuliah,
	})
}

//Create Data
func CreateMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	fmt.Println("masuk setelah db")
	//validasi inputan
	var dataInputMatkul MataKuliahInput
	err := c.ShouldBindJSON(&dataInputMatkul)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

	} else {
		//proses input data
		matkul := models.MataKuliah{
			KodeMatkul:    dataInputMatkul.KodeMatkul,
			NamaMatkul:    dataInputMatkul.NamaMatkul,
			JumlahSKS:     dataInputMatkul.JumlahSKS,
			DosenPengampu: dataInputMatkul.DosenPengampu,
		}

		db.Create(&matkul) //Create DB MySQL

		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully insert data",
			"Data":    matkul,
		})
	}
}
