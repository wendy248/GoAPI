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

type MahasiswaUpdate struct {
	Nama          string `json:"nama" binding:"min=6"`
	Prodi         string `json:"prodi"`
	Fakultas      string `json:"fakultas"`
	NIM           int64  `json:"nim" binding:"numeric,min=100000"`
	TahunAngkatan int64  `json:"tahun" binding:"numeric"`
}

//Read Data
func ReadData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mhs []models.Mahasiswa
	db.Find(&mhs)

	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
	})
}

//Create Data
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

//Update Data
func UpdateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi data
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data mahasiswa tidak di temukan",
		})
		return
	}

	//validasi inputan
	var dataInput MahasiswaUpdate
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	proses Ubah data
	db.Model(&mhs).Update(&dataInput)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfull to Update Data",
		"Data":    mhs,
	})
}

// Delete Data
func DeleteData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Query("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete data",
		})
		return
	}

	db.Delete(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"Data": "Success to delete data",
	})
}
