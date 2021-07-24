package controller

import (
	"net/http"

	"gorm-model/model"
	"gorm-model/util"

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
	var users []model.User

	db := c.MustGet("db").(*gorm.DB)
	db.Find(&users)

	if users == nil {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	responses := make([]model.Response, len(users))

	for i, user := range users {
		responses[i] = model.Response{CreatedAt: util.ReformatDate(user.CreatedAt), UpdatedAt: util.ReformatDate(user.UpdatedAt), UserID: user.ID, UserName: user.Name}
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(responses),
		"users": responses,
	})
}

// GetCompanies method:GET, endpoint:/companies
func GetCompanies(c *gin.Context) {
	var companies []model.Company

	db := c.MustGet("db").(*gorm.DB)
	db.Find(&companies)

	if companies == nil {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	responses := make([]model.Response, len(companies))

	for i, company := range companies {
		responses[i] = model.Response{CompanyID: company.ID, CompanyName: company.Name}
	}

	c.JSON(http.StatusOK, gin.H{
		"count":     len(responses),
		"companies": responses,
	})
}

// GetUserCompanies method:GET, endpoint:/user-companies
func GetUserCompanies(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []model.User

	db.Preload("Company").Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"count": len(users),
		"data":  users,
	})
}

// GetUserCompany method:GET, endpoint:/user-company/:id
func GetUserCompany(c *gin.Context) {
	var user model.User
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Preload("Company").First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
