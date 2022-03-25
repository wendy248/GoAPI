package controller

import (
	"GoAPI/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	Nama          string `json:"nama" binding:"required,min=6"`
	Prodi         string `json:"prodi" binding:"required"`
	Fakultas      string `json:"fakultas" binding:"required"`
	NIM           int64  `json:"nim" binding:"required,numeric,min=100000"`
	TahunAngkatan int16  `json:"tahun" binding:"required,numeric"`
}

//Read Data
func ReadData(c *gin.Context) {
	db := c.MustGet("test").(*gorm.DB)
	var mhs []models.Mahasiswa
	db.Find(&mhs)

	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
	})
}

//POST data >> Create Data
func CreateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi inputan
	var dataInput MahasiswaInput
	err := c.ShouldBindJSON(&dataInput)

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
		mhs := models.Mahasiswa{
			Nama:          dataInput.Nama,
			Prodi:         dataInput.Prodi,
			Fakultas:      dataInput.Fakultas,
			NIM:           dataInput.NIM,
			TahunAngkatan: dataInput.TahunAngkatan,
		}

		db.Create(&mhs) //Create DB MySQL
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully insert data",
			"Data":    mhs,
		})
	}
}
