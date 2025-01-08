package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/octagon")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Set connection pool settings
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(25)
	DB.SetConnMaxLifetime(0)
}
