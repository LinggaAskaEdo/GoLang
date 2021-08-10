package service

import (
	dto "gorm-model/model/dto"
	entity "gorm-model/model/entity"
	"gorm-model/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
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

// Logout function
func Logout(context *gin.Context) {
	au, err := util.ExtractTokenMetadata(context.Request)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error()})
		return
	}

	deleted, delErr := util.DeleteAuth(context, au.AccessUUID)

	if delErr != nil || deleted == 0 { //if any goes wrong
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully logged out"})
}

// Refresh function
func Refresh(context *gin.Context, token *jwt.Token, err error) {
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims

	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string) //convert the interface to string

		if !ok {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  err})
			return
		}

		userID, ok := claims["user_id"].(string)

		if !ok {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  "Error occurred"})
			return
		}

		//Get data by Refresh Token
		fetchData, getErr := util.FetchAuth(context, refreshUUID)

		if getErr != nil || fetchData == 0 { //if any goes wrong
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Unauthorized"})
			return
		}

		//Delete the previous Refresh Token
		deleted, delErr := util.DeleteAuth(context, refreshUUID)

		if delErr != nil || deleted == 0 { //if any goes wrong
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Unauthorized"})
			return
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := util.CreateToken(userID)

		if createErr != nil {
			context.JSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  createErr.Error()})
			return
		}

		//save the tokens metadata to redis
		saveErr := util.CreateAuth(context, fetchData, ts)

		if saveErr != nil {
			context.JSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  saveErr.Error()})
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		context.JSON(http.StatusCreated, tokens)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Refresh expired"})
	}
}
