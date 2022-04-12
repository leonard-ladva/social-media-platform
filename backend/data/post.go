package data

import (
	"time"

	"git.01.kood.tech/Rostislav/real-time-forum/errors"
	uuid "github.com/satori/go.uuid"
)

func (post *Post) Insert() {
	stmt, err := DB.Prepare("INSERT INTO Post (ID, UserID, Content, TagID, CreatedAt) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		errors.ErrorLog.Print("Database: failed when inserting post to database.")
	}
	defer stmt.Close()

	post.TagID = getTagId(post.Tag)

	id := uuid.NewV4()
	createdAt := time.Now().UnixNano()
	// temporary
	post.UserID = "beb8e5c6-c33e-4306-8725-fa07aa89f2f4"
	stmt.Exec(id, post.UserID, post.Content, post.TagID, createdAt)
}

func getTagId(title string) string {
	var post Post
	query := "SELECT ID FROM Tag WHERE Title = ?"
	row := DB.QueryRow(query, title)
	err := row.Scan(&post.TagID)
	if err != nil {
		errors.ErrorLog.Print(err)
	}

	return post.TagID
}
