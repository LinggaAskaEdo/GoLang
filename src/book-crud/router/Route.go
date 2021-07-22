package router

import (
	"book-crud/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/tasks", controller.FindTasks)
	r.POST("/tasks", controller.CreateTask)
	r.GET("/tasks/:id", controller.FindTask)
	r.PATCH("/tasks/:id", controller.UpdateTask)
	r.DELETE("tasks/:id", controller.DeleteTask)

	return r
}
