package data

import (
	"database/sql"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

var postAllowedChars = map[string]string{
	"Content": "^[a-zA-Z0-9]*$", // letters, numbers
	"Tag":     "^[ -~]*$",       // all ascii characres in range space to ~
}

func (post *Post) Insert() error {
	stmt, err := DB.Prepare("INSERT INTO Post (ID, UserID, Content, TagID, CreatedAt) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("data: inserting post")
	}
	defer stmt.Close()

	tag, err := getTagByTitle(post.Tag)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("data: getting post TagID")
	}
	if err != nil {
		tag = &Tag{}
		tag.Title = post.Tag
		tag, err = tag.Insert()
		if err != nil {
			return err
		}
	}
	post.TagID = tag.ID
	id := uuid.NewV4()
	createdAt := CurrentTime()

	stmt.Exec(id, post.UserID, post.Content, post.TagID, createdAt)
	return nil
}

func LatestPosts(lastEarliestPost string) ([]*Post, error) {
	var posts []*Post
	query := "SELECT ID, Content, TagID, UserId, CreatedAt FROM Post ORDER BY CreatedAt DESC LIMIT 5"
	if lastEarliestPost != "-1" {
		query = fmt.Sprintf("SELECT ID, Content, TagID, UserId, CreatedAt FROM Post WHERE CreatedAt < %s ORDER BY CreatedAt DESC LIMIT 5", lastEarliestPost)
	}
	rows, err := DB.Query(query)
	if err == sql.ErrNoRows {
		return posts, nil
	}
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("data: getting posts")
	}

	for rows.Next() {
		post := &Post{}

		err := rows.Scan(&post.ID, &post.Content, &post.TagID, &post.UserID, &post.CreatedAt)
		if err != nil {
			return nil, errors.New("data: getting posts")
		}
		posts = append(posts, post)
	}

	return posts, nil
}
