package database

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func AddSession(w http.ResponseWriter, r *http.Request, user User) error {

	log.Println("AddSession, user: ", user.Username)
	Db.Exec("DELETE FROM session WHERE user_id = $1", user.ID)
	log.Println("Adding session to database for user: ", user.ID, user.Username)

	sessionID := uuid.New()
	cookie := &http.Cookie{
		Name:   "session",
		Value:  sessionID.String(),
		Secure: true,
		Path:   "/",
	}
	cookie.MaxAge = 60 * 60 // 5 minutes (for testing, increase later)
	http.SetCookie(w, cookie)

	statement, err := Db.Prepare("INSERT INTO session (user_id, uuid, datetime) VALUES (?, ?, ?)")
	defer statement.Close()
	if err != nil {
		log.Println("Error adding session to database (session.go)")
		return err
	}

	statement.Exec(user.ID, sessionID, time.Now().Add(60*time.Minute)) //cookie expiry time 5 min
	http.SetCookie(w, cookie)

	return nil
}
