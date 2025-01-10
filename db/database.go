package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}

	dbUser := os.Getenv("DB_USER_MASTER")
	dbPass := os.Getenv("DB_PASS_MASTER")
	dbHost := os.Getenv("DB_PASS_MASTER")
	dbPort := os.Getenv("DB_PORT_MASTER")
	dbName := os.Getenv("DB_NAME_MASTER")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Set connection pool settings
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(25)
	DB.SetConnMaxLifetime(0)
}
