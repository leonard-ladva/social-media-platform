package data

import (
	"database/sql"
	"errors"
)

func (chat Chat) Insert() error {

	stmt, err := DB.Prepare("INSERT INTO Chat (ID, LastMessageTime, CreatedAt) VALUES (?, ?, ?)")
	if err != nil {
		return errors.New("data: inserting chat")
	}
	defer stmt.Close()

	chat.CreatedAt = CurrentTime()

	stmt.Exec(chat.ID, chat.LastMessageTime, chat.CreatedAt)
	return nil
}

func (chat Chat) Update() error {
	stmt, err := DB.Prepare(`UPDATE Chat SET LastMessageTime = ? WHERE ID = "?"`)
	if err != nil {
		return errors.New("data: updating chat")
	}
	defer stmt.Close()

	stmt.Exec(chat.LastMessageTime, chat.ID)
	return nil
}

func (chat Chat) Get() (Chat, error) {
	query := "SELECT ID, LastMessageTime, CreatedAt FROM Chat WHERE ID = ?"
	row := DB.QueryRow(query, chat.ID)

	err := row.Scan(&chat.ID, &chat.LastMessageTime, &chat.CreatedAt)
	if err != nil {
		return chat, err
	}

	return chat, nil
}

func (chat Chat) Exists() (bool, error) {
	chat, err := chat.Get()
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
