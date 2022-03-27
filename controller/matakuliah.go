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
	KodeMatkul    string `json:"kode matkul" binding:"required"`
	NamaMatkul    string `json:"nama matkul" binding:"required,min=4"`
	JumlahSKS     int16  `json:"jumlah sks" binding:"required"`
	DosenPengampu string `json:"dosen pengampu" binding:"required"`
}

type MataKuliahUpdate struct {
	KodeMatkul    string `json:"kode matkul"`
	NamaMatkul    string `json:"nama matkul" binding:"min=4"`
	JumlahSKS     int16  `json:"jumlah sks"`
	DosenPengampu string `json:"dosen pengampu"`
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
			"error": err.Error(),
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

//Update Data
func UpdateMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi data
	var matakuliah models.MataKuliah
	if err := db.Where("kode_matkul = ?", c.Param("kode")).First(&matakuliah).Error; 
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data Mata Kuliah tidak di temukan",
		})
		return
	}

	// if err2 := db.Where("kode_matku")

	//validasi inputan
	var dataInput MataKuliahUpdate
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	proses Ubah data
	db.Model(&matakuliah).Update(&dataInput)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfull to Update Data",
		"Data":    matakuliah,
	})
}

// Delete Data
func DeleteMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var matakuliah models.MataKuliah
	if err := db.Where("kode_matkul = ?", c.Query("kode")).First(&matakuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not found in database",
		})
		return
	}

	db.Delete(&matakuliah)
	c.JSON(http.StatusOK, gin.H{
		"Data": "Success to delete data",
	})
}
