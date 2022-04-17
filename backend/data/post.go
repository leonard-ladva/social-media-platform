package data

import (
	"time"
	"errors"

	// "git.01.kood.tech/Rostislav/real-time-forum/errors"
	uuid "github.com/satori/go.uuid"
)

func (post *Post) Insert() error {
	stmt, err := DB.Prepare("INSERT INTO Post (ID, UserID, Content, TagID, CreatedAt) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		return errors.New("data: inserting post")
	}
	defer stmt.Close()

	post.TagID, err = getTagId(post.Tag)
	if err != nil {
		return errors.New("data: getting post TagID")
	}
	id := uuid.NewV4()
	createdAt := time.Now().UnixNano()

	// temporary
	post.UserID = "beb8e5c6-c33e-4306-8725-fa07aa89f2f4"
	stmt.Exec(id, post.UserID, post.Content, post.TagID, createdAt)
	return nil
}

// getTagId gets the tagId of the post from the database
func getTagId(title string) (string, error) {
	var post Post
	query := "SELECT ID FROM Tag WHERE Title = ?"
	row := DB.QueryRow(query, title)
	err := row.Scan(&post.TagID)
	if err != nil {
		return "", err
	}
	return post.TagID, nil
}
