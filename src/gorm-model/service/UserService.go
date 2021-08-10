package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	dto "gorm-model/model/dto"
	entity "gorm-model/model/entity"
	"gorm-model/util"
)

// Login function
func Login(context *gin.Context, input dto.Auth) {
	var cred entity.Credential

	db := context.MustGet("db").(*gorm.DB)
	db.Limit(1).Find(&cred)

	if &cred == nil {
		context.JSON(http.StatusNoContent, nil)
		return
	}

	if input.Email != cred.Email || input.Password != cred.Password {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Please provide valid login details"})
		return
	}

	token, err := util.CreateToken(cred.Email)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveToken := util.CreateAuth(context, cred.ID, token)

	if saveToken != nil {
		context.JSON(http.StatusUnprocessableEntity, saveToken.Error())
	}

	context.JSON(http.StatusOK, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken})
}

// CreateUserService function
func CreateUserService(context *gin.Context, input dto.Request) {
	user := entity.User{
		Name:    input.UserName,
		Company: entity.Company{Name: input.CompanyName}}

	db := context.MustGet("db").(*gorm.DB)
	db.Create(&user)

	context.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUsersService function
func GetUsersService(context *gin.Context) {
	var users []entity.User

	db := context.MustGet("db").(*gorm.DB)
	db.Find(&users)

	if users == nil {
		context.JSON(http.StatusNoContent, nil)
		return
	}

	responses := make([]dto.Response, len(users))

	for i, user := range users {
		responses[i] = dto.Response{CreatedAt: util.ReformatDate(user.CreatedAt), UpdatedAt: util.ReformatDate(user.UpdatedAt), UserID: user.ID, UserName: user.Name}
	}

	context.JSON(http.StatusOK, gin.H{
		"count": len(responses),
		"users": responses,
	})
}

// GetCompaniesService function
func GetCompaniesService(context *gin.Context) {
	var companies []entity.Company

	db := context.MustGet("db").(*gorm.DB)
	db.Find(&companies)

	if companies == nil {
		context.JSON(http.StatusNoContent, nil)
		return
	}

	responses := make([]dto.Response, len(companies))

	for i, company := range companies {
		responses[i] = dto.Response{CompanyID: company.ID, CompanyName: company.Name}
	}

	context.JSON(http.StatusOK, gin.H{
		"count":     len(responses),
		"companies": responses,
	})
}

// GetUserCompaniesService function
func GetUserCompaniesService(context *gin.Context) {
	var users []entity.User

	db := context.MustGet("db").(*gorm.DB)
	db.Preload("Company").Find(&users)

	context.JSON(http.StatusOK, gin.H{
		"count": len(users),
		"data":  users,
	})
}

// GetUserCompanyService function
func GetUserCompanyService(context *gin.Context) {
	var user entity.User

	db := context.MustGet("db").(*gorm.DB)

	if err := db.Preload("Company").First(&user, context.Param("id")).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}
