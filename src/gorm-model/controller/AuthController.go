package controller

import (
	"fmt"
	dto "gorm-model/model/dto"
	"gorm-model/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Login method:POST, endpoint:/login
func Login(context *gin.Context) {
	var auth dto.Auth

	if err := context.ShouldBindJSON(&auth); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Invalid json provided"})
		return
	}

	service.Login(context, auth)
}

// Logout method:POST, endpoint:/logout
func Logout(context *gin.Context) {
	service.Logout(context)
}

// Refresh function
func Refresh(context *gin.Context) {
	mapToken := map[string]string{}

	if err := context.ShouldBindJSON(&mapToken); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err.Error()})
		return
	}

	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Refresh token expired"})
		return
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err})
		return
	}

	service.Refresh(context, token, err)
}
