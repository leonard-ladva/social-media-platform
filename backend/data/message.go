package data

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

func (message *Message) Insert() error {

	stmt, err := DB.Prepare("INSERT INTO Message (ID, ChatID, UserID, Content, CreatedAt) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("db: inserting message")
	}
	defer stmt.Close()

	message.ID = uuid.NewV4().String()

	stmt.Exec(message.ID, message.ChatID, message.UserID, message.Content, message.CreatedAt)
	return nil
}

// func (chat *Chat) GetLatest() (*Chat, error) {
// 	query := "SELECT (ID, LastMessageTime, CreatedAt) FROM Chat WHERE ChatID = ?"
// 	row := DB.QueryRow(query, chat.ID)

// 	err := row.Scan(&chat.ID, &chat.LastMessageTime, &chat.CreatedAt)
// 	if err != nil {
// 		return chat, err
// 	}

// 	return chat, nil
// }
