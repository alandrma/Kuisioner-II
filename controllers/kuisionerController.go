package controllers

import (
	"net/http"

	"Kuisioner-MySql/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateKuisioners struct {
	Judul string `json:"JudulKuisioner"`
	Isi   string `json:"IsiKuisioner"`
}

type UpdateKuisioners struct {
	Judul string `json:"JudulKuisioner"`
	Isi   string `json:"IsiKuisioner"`
}

func GetKuis(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var kuisioner []models.Kuisioner
	db.Find(&kuisioner)
	c.JSON(http.StatusOK, gin.H{"data": kuisioner})
}

func TambahKuis(c *gin.Context) {
	var input CreateKuisioners
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kuisioner := models.Kuisioner{JudulKuisioner: input.Judul, IsiKuisioner: input.Isi}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&kuisioner)

	c.JSON(http.StatusOK, gin.H{"data": kuisioner})
}

func CariKuis(c *gin.Context) {
	var kuisioner models.Kuisioner

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&kuisioner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kuisioner})
}

func Search(c *gin.Context) {
	var kuisioner models.Kuisioner

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("isi LIKE ?", c.Param("isi")).Find(&kuisioner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kuisioner})
}

func UpdateKuis(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var kuisioner models.Kuisioner
	if err := db.Where("id = ?", c.Param("id")).First(&kuisioner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	var input UpdateKuisioners
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Kuisioner
	updatedInput.JudulKuisioner = input.Judul
	updatedInput.IsiKuisioner = input.Isi

	db.Model(&kuisioner).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": kuisioner})
}

func DeleteKuis(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var kuisioner models.Kuisioner
	if err := db.Where("id = ?", c.Param("id")).First(&kuisioner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&kuisioner)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
