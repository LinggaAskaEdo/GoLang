package main

import (
	"book-crud/config"
	"book-crud/model"
	"book-crud/router"
)

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&model.Task{})

	r := router.SetupRoutes(db)
	r.Run()
}
