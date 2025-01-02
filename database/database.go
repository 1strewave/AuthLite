package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Initialize() {
	var err error
	DB, err = sql.Open("sqlite3", "./authlite.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	log.Println("Database successfully connected!")
}
