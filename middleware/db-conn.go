package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3" // import sqlite package
)

func init() {
	// Load env variables using the godotenv package
	// DB details is saved as env variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// open the db connection
	database, err := sql.Open("sqlite3", dbName)

	if err != nil {
		log.Fatalf("Error: Sqlite connection")
	}

	// create mac table query
	createTableStatement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (mac_id INTEGER PRIMARY KEY, mac_address TEXT NOT NULL UNIQUE, mac_group TEXT NOT NULL)", dbName)

	// run the create table query
	statement, _ := database.Prepare(createTableStatement)

	statement.Exec()
}

func getDBConn() *sql.DB {
	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// open sqlite db connection
	database, err := sql.Open("sqlite3", dbName)

	if err != nil {
		log.Fatalf("Error: Sqlite connection")
	}

	fmt.Println("DB Connected ")
	return database
}
