package config

import (
	"fmt"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupDB : initializing mysql database connection
func SetupDB() *gorm.DB {
	USER := "lingga.putra"
	PASS := "a5k4CooL"
	HOST := "127.0.0.1"
	PORT := "3306"
	DBNAME := "GORM-MODEL"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}

	return db
}
