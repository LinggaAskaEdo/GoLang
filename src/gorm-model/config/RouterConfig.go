package config

import (
	"gorm-model/controller"

	"github.com/go-redis/redis/v7"

	"gorm-model/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRoutes function
func SetupRoutes(db *gorm.DB, redis *redis.Client) *gin.Engine {
	router := gin.Default()

	router.Use(func(context *gin.Context) {
		context.Set("db", db)
		context.Set("redis", redis)
	})

	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)
	router.POST("/refresh", controller.Refresh)

	// ONE to ONE relationship
	router.POST("/user", util.TokenAuthMiddleware(), controller.CreateUser)
	router.GET("/users", util.TokenAuthMiddleware(), controller.GetUsers)
	router.GET("/companies", util.TokenAuthMiddleware(), controller.GetCompanies)
	router.GET("/user-companies", util.TokenAuthMiddleware(), controller.GetUserCompanies)
	router.GET("/user-company/:id", util.TokenAuthMiddleware(), controller.GetUserCompany)

	return router
}
