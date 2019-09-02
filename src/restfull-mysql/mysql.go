package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:L1nggaa5k43d0@tcp(127.0.0.1:3306)/GO_TEST")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
