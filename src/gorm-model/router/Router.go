package router

import (
	"gorm-model/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRoutes function
func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/user", controller.CreateUser)
	r.GET("/users", controller.GetUsers)
	r.GET("/user/:id", controller.GetUser)

	return r
}
