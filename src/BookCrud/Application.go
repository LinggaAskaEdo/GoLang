package main

import (
	"BookCrud/config"
	"BookCrud/model"
	"BookCrud/router"
)

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&model.Task{})

	r := router.SetupRoutes(db)
	r.Run()
}
