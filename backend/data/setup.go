package data

import (
	"database/sql"
	"errors"


	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Connect opens a sql database file
func Connect() error {
	db, err := sql.Open("sqlite3", "file:database.db")
	if err != nil {
		return errors.New("data: connecting to database")
	}
	err = db.Ping()
	if err != nil {
		return errors.New("data: pinging database")
	}
	DB = db
	return nil
}
