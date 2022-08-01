package controllers

import (
	"net/http"

	"Kuisioner-MySql/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUsers struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	UserName     string `json:"UserName"`
	Email        string `json:"Email"`
	PasswordHash string `json:"PasswordHash"`
}

type UpdateUsers struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	UserName     string `json:"UserName"`
	Email        string `json:"Email"`
	PasswordHash string `json:"PasswordHash"`
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func TambahUser(c *gin.Context) {
	var input CreateUsers
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{FirstName: input.FirstName, LastName: input.LastName, UserName: input.UserName, Email: input.Email, PasswordHash: input.PasswordHash}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CariUser(c *gin.Context) {
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	var input UpdateUsers
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.User
	updatedInput.FirstName = input.FirstName
	updatedInput.LastName = input.LastName
	updatedInput.UserName = input.UserName
	updatedInput.Email = input.Email
	updatedInput.PasswordHash = input.PasswordHash

	db.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
