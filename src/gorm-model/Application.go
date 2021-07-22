package main

import (
	"gorm-model/config"
	"gorm-model/model"
	"gorm-model/router"
)

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&model.User{}, &model.Company{})

	r := router.SetupRoutes(db)
	r.Run()
}
