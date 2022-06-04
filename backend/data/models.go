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
	ID              string `json:"id"`
	LastMessageTime int64 `json:"lastMessageTime"`
	CreatedAt       int64 `json:"createdAt"`
}

type Comment struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	PostID    string `json:"postId"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
}

type Message struct {
	ID         string `json:"id"`
	ChatID     string `json:"chatId"`
	UserID     string `json:"userId"`
	ReceiverID string `json:"receiverId"`
	Content    string `json:"content"`
	CreatedAt  int64  `json:"createdAt"`
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
