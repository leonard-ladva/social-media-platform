package data

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

// AddSession adds a new session to the database
func (user User) AddSession() (uuid.UUID, error) {
	DB.Exec("DELETE FROM Session WHERE UserID = $1", user.ID)
	token := uuid.NewV4()

	statement, err := DB.Prepare("INSERT INTO Session (ID, UserID, CreatedAt) VALUES (?, ?, ?)")
	defer statement.Close()
	if err != nil {
		return token, errors.New("data: inserting session")
	}

	statement.Exec(token, user.ID, CurrentTime())
	return token, nil
}

// GetSession gets a session from the database
func GetSession(token string) (Session, error) {
	row := DB.QueryRow("SELECT ID, UserID, CreatedAt FROM Session WHERE ID = ?", token)
	var session Session
	err := row.Scan(&session.ID, &session.UserID, &session.CreatedAt)
	if err != nil {
		return session, errors.New("data: getting session")
	}
	return session, nil
}
