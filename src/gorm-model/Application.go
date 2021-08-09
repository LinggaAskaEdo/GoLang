package main

import (
	"gorm-model/config"
	entity "gorm-model/model/entity"
)

func main() {
	// Initializing database
	db := config.SetupDB()
	db.AutoMigrate(&entity.User{}, &entity.Company{}, &entity.Credential{})

	// Initializing redis
	redisClient := config.SetupRedis()

	router := config.SetupRoutes(db, redisClient)
	err := router.Run()

	if err != nil {
		return
	}
}
