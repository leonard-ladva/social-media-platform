package data

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	uuid "github.com/satori/go.uuid"
)

func (comment *Comment) Insert() error {
	stmt, err := DB.Prepare("INSERT INTO Comment (ID, UserID, PostID, Content, CreatedAt) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("data: inserting comment")
	}
	defer stmt.Close()

	id := uuid.NewV4()
	createdAt := CurrentTime()

	stmt.Exec(id, comment.UserID, comment.PostID, comment.Content, createdAt)
	return nil
}

func LatestComments(lastEarliestComment string) ([]*Comment, error) {
	var comments []*Comment
	query := "SELECT ID, UserID, PostID, Content, CreatedAt FROM Comment ORDER BY CreatedAt DESC LIMIT 10"
	if lastEarliestComment != "-1" {
		query = fmt.Sprintf("SELECT ID, UserID, PostID, Content, CreatedAt FROM Comment WHERE CreatedAt < %s ORDER BY CreatedAt DESC LIMIT 10", lastEarliestComment)
	}
	rows, err := DB.Query(query)
	if err == sql.ErrNoRows {
		return comments, nil
	}
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("data: getting comments")
	}

	for rows.Next() {
		comment := &Comment{}

		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt)
		if err != nil {
			return nil, errors.New("data: getting comments")
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

// IsValid checks all the comment fields and returns true if valid
func (c *Comment) IsValid() (bool, error) {
	errs := url.Values{}
	// Content
	if !checkCharacters("Content", c.Content) || !checkLength("Content", c.Content) {
		errs.Add("Content", fmt.Sprintf("Content has to be between %d and %d characters",
			lengthReq["Content"][0], lengthReq["Content"][1]))
	}
	if len(errs) != 0 {
		fmt.Println("Form Errors: ", errs)
		return false, nil
	}
	return true, nil
}
