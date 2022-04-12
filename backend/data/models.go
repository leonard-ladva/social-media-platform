package data

import (
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

type Category struct {
	ID    string
	Title string
}

type Chat struct {
	ID           string
	MessageCount int
}

type Comment struct {
	ID        string
	UserID    string
	PostID    string
	Content   string
	CreatedAt time.Time
}

type Message struct {
	ChatID    string
	UserID    string
	MessageNr int
	Content   string
	CreatedAt time.Time
}

type Post struct {
	ID			string		`json:"id"`
	UserID	    string		`json:"userId"`
	Content		string		`json:"content"`
	Tag			string		`json:"tag"`
	TagID		string		`json:"tagId"`
	CreatedAt	time.Time	`json:"createdAt"`
}

type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
}

type User struct {
	ID				string		`json:"id"`
	Email			string		`json:"email"`
	Password		[]byte		`json:"password"`
	PasswordPlain	string		`json:"passwordPlain"`
	PasswordConfirm	string		`json:"passwordConfirm"`
	Nickname		string		`json:"nickname"`
	FirstName		string		`json:"firstName"`
	LastName		string		`json:"lastName"`
	Gender			string		`json:"gender"`
	Age				int			`json:"age"`
	Color			string		`json:"color"`
	CreatedAt		time.Time	`json:"createdAt"`
}
