package data

import (
	"database/sql"
	"errors"

	uuid "github.com/satori/go.uuid"
)

// getTagId gets the tagId of the post from the database
func getTagByTitle(title string) (*Tag, error) {
	tag := &Tag{}
	query := "SELECT ID, Title, CreatedAt FROM Tag WHERE Title = ?"
	row := DB.QueryRow(query, title)
	err := row.Scan(&tag.ID, &tag.Title, &tag.CreatedAt)
	if err != nil {
		return nil, errors.New("data: getting tag") 
	}
	return tag, nil
}
// GetTagByID gets a tag from the database given an ID
func GetTagByID(ID string) (*Tag, error) {
	tag := &Tag{}
	query := "SELECT ID, Title, CreatedAt FROM Tag WHERE ID = ?"
	row := DB.QueryRow(query, ID)
	err := row.Scan(&tag.ID, &tag.Title, &tag.CreatedAt)
	if err != nil {
		return nil, errors.New("data: getting tag")
	}
	return tag, nil
}

// Insert inserts a new tag into the database
func (tag *Tag) Insert() (*Tag, error) {
	stmt, err := DB.Prepare("INSERT INTO Tag (ID, Title, CreatedAt) VALUES (?, ?, ?);")
	if err != nil {
		return nil, errors.New("data: inserting tag")
	}
	defer stmt.Close()

	tag.ID = uuid.NewV4().String()
	tag.CreatedAt = currentTime()

	stmt.Exec(tag.ID, tag.Title, tag.CreatedAt)
	return tag, nil
}

func GetAllTags() ([]*Tag, error) {
	var tags []*Tag
	query := "SELECT * FROM Tag"
	rows, err := DB.Query(query)
	if err == sql.ErrNoRows {
		return tags, nil
	}
	if err != nil {
		return tags, errors.New("data: getting all tags")
	}

	for rows.Next() {
		tag := &Tag{}

		err := rows.Scan(&tag.ID, &tag.Title, &tag.CreatedAt)
		if err != nil {
			return nil, errors.New("data: getting all tags")
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
