package data

import (
	"database/sql"
)

type DBModel struct {
	DB *sql.DB
}

type Tag struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"createdAt"`
}

type Chat struct {
	ID              string
	LastMessageTime int64
	CreatedAt       int64
}

type Comment struct {
	ID        string
	UserID    string
	PostID    string
	Content   string
	CreatedAt int64
}

type Message struct {
	ID        string
	ChatID    string
	UserID    string
	Content   string
	CreatedAt int64
}

type Post struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Content   string `json:"content"`
	Tag       string `json:"tag"`
	TagID     string `json:"tagId"`
	CreatedAt int64  `json:"createdAt"`
}

type Session struct {
	ID        string
	UserID    string
	CreatedAt int64
}

type User struct {
	ID              string    `json:"id"`
	Email           string    `json:"email"`
	Password        []byte    `json:"password"`
	PasswordPlain   string    `json:"passwordPlain"`
	PasswordConfirm string    `json:"passwordConfirm"`
	Nickname        string    `json:"nickname"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Gender          string    `json:"gender"`
	Age             StringInt `json:"age"`
	Color           string    `json:"color"`
	CreatedAt       int64     `json:"createdAt"`
	Active          bool      `json:"active"`
}
