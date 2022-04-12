package data

import (
	"database/sql"

	"git.01.kood.tech/Rostislav/real-time-forum/errors"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open("sqlite3", "file:database.db")
	if err != nil {
		errors.ErrorLog.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		errors.ErrorLog.Fatal(err)
	}
	DB = db
}
