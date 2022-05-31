package data

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

func (message Message) Insert() error {
	stmt, err := DB.Prepare("INSERT INTO Message (ID, ChatID, UserID, Content, CreatedAt) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("db: inserting message")
	}
	defer stmt.Close()

	message.ID = uuid.NewV4().String()

	stmt.Exec(message.ID, message.ChatID, message.UserID, message.Content, message.CreatedAt)
	return nil
}

func GetLatestMessages(lastEarliestMessage string, chatID string) (messages []*Message, err error) {
	query := "SELECT ID, ChatID, UserID, Content, CreatedAt FROM Message WHERE ChatID = ? AND CreatedAt < ? ORDER BY CreatedAt DESC LIMIT 20"
	rows, err := DB.Query(query, chatID, lastEarliestMessage)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		message := &Message{}
		err := rows.Scan(&message.ID, &message.ChatID, &message.UserID, &message.Content, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (message Message) Valid() (bool, error) {
	senderExsists, _, err := GetUser("ID", message.UserID)
	if err != nil {
		return false, err
	}

	receiverExsists, _, err := GetUser("ID", message.ReceiverID)
	if err != nil {
		return false, err
	}

	if !senderExsists || !receiverExsists || message.Content == "" {
		return false, nil
	}
	return true, nil
}