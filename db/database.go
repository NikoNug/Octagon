package db

import (
	"database/sql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/octagon")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
