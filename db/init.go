package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	// Open a connection to the SQLite database
	db, err = sql.Open("sqlite3", "db/app.db")
	if err != nil {
		return nil, err
	}

	// Create a table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE
	);`
	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, err
	}

	return db, nil
}
