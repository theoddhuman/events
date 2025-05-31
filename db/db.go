package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "./db.sqlite3")

	if err != nil {
		panic("Could not connect to database")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}
