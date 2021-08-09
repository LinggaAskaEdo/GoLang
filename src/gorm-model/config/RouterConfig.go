package config

import (
	"gorm-model/controller"

	"github.com/go-redis/redis/v7"

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

	// ONE to ONE relationship
	router.POST("/user", controller.CreateUser)
	router.GET("/users", controller.GetUsers)
	router.GET("/companies", controller.GetCompanies)
	router.GET("/user-companies", controller.GetUserCompanies)
	router.GET("/user-company/:id", controller.GetUserCompany)

	return router
}
