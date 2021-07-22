package controller

import (
	"net/http"

	"gorm-model/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateUser method:POST, endpoint:/user
func CreateUser(c *gin.Context) {
	// Validate input
	var input model.Request

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User
	user := model.User{
		Name:    input.UserName,
		Company: model.Company{Name: input.CompanyName}}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUsers method:GET, endpoint:/users
func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []model.User

	// db.Find(&users)
	db.Preload("Company").Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"count": len(users),
		"data":  users,
	})
}

// GetUser method:GET, endpoint:/user/:id
func GetUser(c *gin.Context) {
	var user model.User
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Preload("Company").First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
