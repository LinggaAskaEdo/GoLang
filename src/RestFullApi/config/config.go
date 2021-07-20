package config

import (
	"RestFullApi/structs"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "lingga.putra:a5k4CooL@tcp(127.0.0.1:3306)/GO_DB?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})

	// sqlDB, err := db.DB()

	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDB.SetMaxOpenConns(100)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
