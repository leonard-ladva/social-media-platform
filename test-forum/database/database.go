package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Env holds environment variables
type Env struct {
	Port 		string
	HTTPSport 	string
	SQLuser 	string
	SQLpass 	string
	DBpath		string
}

var Db *sql.DB

func Initialize(env *Env) {

	dbTables := []string{
		`CREATE TABLE IF NOT EXISTS "role" (
			"role_id"			INTEGER UNIQUE NOT NULL,
			"role"				TEXT UNIQUE NOT NULL,
			PRIMARY KEY("role_id" AUTOINCREMENT)
		)`,

		`CREATE TABLE IF NOT EXISTS "user" (
			"user_id"			INTEGER UNIQUE NOT NULL,
			"username"			TEXT UNIQUE NOT NULL,
			"password"			TEXT NOT NULL,
			"email"				TEXT UNIQUE NOT NULL, 
			"reg_datetime"		DATETIME DEFAULT CURRENT_TIMESTAMP,
			"role_id"			INTEGER NOT NULL,
			PRIMARY KEY("user_id" AUTOINCREMENT)
			FOREIGN KEY("role_id") REFERENCES "ROLE"("role_id")
		)`,

		`CREATE TABLE IF NOT EXISTS "notification" (
			"notification_id"	INTEGER UNIQUE NOT NULL,
			"user_id"			INTEGER NOT NULL,
			"content"			TEXT NOT NULL,
			"priority"			INTEGER NOT NULL,
			"datetime"			DATETIME DEFAULT CURRENT_TIMESTAMP,
			"read"				INTEGER NOT NULL,
			"link"				TEXT NOT NULL,
			PRIMARY KEY("notification_id" AUTOINCREMENT)
			FOREIGN KEY("user_id") REFERENCES "USER"("user_id")
		)`,

		`CREATE TABLE IF NOT EXISTS "post" (
			"post_id"	INTEGER NOT NULL UNIQUE,
			"title"		TEXT NOT NULL,
			"content"	TEXT NOT NULL,
			"user_id"	INTEGER NOT NULL,
			"datetime"	DATETIME DEFAULT CURRENT_TIMESTAMP,
			"postimg" 	TEXT,
			"status"	TEXT NOT NULL,
			PRIMARY KEY("post_id" AUTOINCREMENT),
			FOREIGN KEY("user_id") REFERENCES "USER"("user_id")
		)`,

		`CREATE TABLE IF NOT EXISTS "comment" (
			"comment_id"	INTEGER NOT NULL UNIQUE,
			"post_id"	INTEGER NOT NULL,
			"user_id"	INTEGER NOT NULL,
			"content"	TEXT NOT NULL,
			"datetime"	DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY("comment_id" AUTOINCREMENT),
			FOREIGN KEY("post_id") REFERENCES "POST"("post_id"),
			FOREIGN KEY("user_id") REFERENCES "USER"("user_id")
		)`,

		`CREATE TABLE IF NOT EXISTS "session" (
			"session_id"	INTEGER NOT NULL UNIQUE,
			"datetime"	DATETIME DEFAULT CURRENT_TIMESTAMP,
			"user_id"	INTEGER NOT NULL,
			"uuid"	TEXT NOT NULL,
			PRIMARY KEY("session_id" AUTOINCREMENT),
			FOREIGN KEY("user_id") REFERENCES "USER"("user_id")
		)`,

		`CREATE TABLE IF NOT EXISTS "category" (
			"category_id"	INTEGER NOT NULL UNIQUE,
			"title"	TEXT NOT NULL,
			"description"	TEXT NOT NULL,
			"img_link" TEXT NOT NULL,
			PRIMARY KEY("category_id" AUTOINCREMENT)
		)`,

		`CREATE TABLE IF NOT EXISTS "post_reaction" (
			"reaction_id"	INTEGER NOT NULL UNIQUE,
			"post_id"	INTEGER NOT NULL,
			"user_id"	INTEGER NOT NULL,
			"type"	TEXT NOT NULL,
			"datetime"	DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY("reaction_id" AUTOINCREMENT),
			FOREIGN KEY("post_id") REFERENCES "POST"("post_id"),
			FOREIGN KEY("user_id") REFERENCES "USER"("user_id")
		)`,

		`CREATE TABLE IF NOT EXISTS "comment_reaction" (
			"reaction_id"	INTEGER NOT NULL UNIQUE,
			"comment_id"	INTEGER NOT NULL,
			"user_id"	INTEGER NOT NULL,
			"type"	TEXT NOT NULL,
			"datetime"	DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY("reaction_id" AUTOINCREMENT),
			FOREIGN KEY("user_id") REFERENCES "USER"("user_id"),
			FOREIGN KEY("comment_id") REFERENCES "COMMENT"("comment_id")
		)`,

		`CREATE TABLE IF NOT EXISTS "post_category" (
			"postcat_id"	INTEGER NOT NULL UNIQUE,
			"post_id"	INTEGER NOT NULL,
			"category_id"	INTEGER NOT NULL,
			PRIMARY KEY("postcat_id" AUTOINCREMENT),
			FOREIGN KEY("post_id") REFERENCES "POST"("post_id"),
			FOREIGN KEY("category_id") REFERENCES "CATEGORY"("category_id")
		)`,
	}

	var err error
	connectionString := env.DBpath + "?_auth&_auth_user=" + env.SQLuser + "&_auth_pass=" + env.SQLpass + "&_auth_crypt=sha1"
	Db, err = sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Println("ERROR | Unable to create database.db")
		panic(err)
	}

	// Create each database table :
	for _, table := range dbTables {
		err := createDatabase(table)
		if err != nil {
			panic(err)
		}
	}
	log.Println("DATABASE | Database created and initialized successfully.")
}

func createDatabase(table string) error {
	statement, err := Db.Prepare(table)
	if err != nil {
		log.Println("ERROR | Unable to create tables in the database.")
		return err
	}

	defer statement.Close()

	statement.Exec()
	return nil
}
