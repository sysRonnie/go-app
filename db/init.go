package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Create the db variable
var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	// Open a connection to the SQLite database
	db, err = sql.Open("sqlite3", "db/app.db")
	if err != nil {
		return nil, err
	}

	// Create a table if it doesn't exist
	sqlCreateTableUSER_MASTER := `
	  CREATE TABLE IF NOT EXISTS USER_MASTER (
		USER_ID INTEGER PRIMARY KEY AUTOINCREMENT,
		USERNAME TEXT NOT NULL UNIQUE,
		EMAIL TEXT NOT NULL UNIQUE,
		DEVICE_NAME TEXT NOT NULL,
		BROWSER_TYPE TEXT NOT NULL,
		IP_ADDRESS TEXT NOT NULL
	);`

	if _, err := db.Exec(sqlCreateTableUSER_MASTER); err != nil {
		return nil, err
	}

	return db, nil
}
