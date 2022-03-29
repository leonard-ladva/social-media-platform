package database

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetUserByUserName(username string) User {

	var user User
	row := Db.QueryRow("SELECT u.user_id, u.username, u.password, u.email, u.reg_datetime, u.role_id, r.role FROM user u LEFT JOIN role r ON r.role_id = u.role_id WHERE u.username = $1 OR u.email = $1", username)

	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Date, &user.RoleID, &user.Role)
	log.Println("Getting user from database | User: ", username)
	return user
}

func GetUserByCookie(w http.ResponseWriter, r *http.Request) User {
	// Get cookie from client
	userCookie, err := r.Cookie("session")
	// If no cookie, create it
	if err != nil {
		sessionID := uuid.New()
		userCookie = &http.Cookie{
			Name:   "session",
			Value:  sessionID.String(),
			Secure: true,
		}
		userCookie.MaxAge = 60 * 60
		http.SetCookie(w, userCookie)
	}
	session := GetSessionByUUID(userCookie.Value)
	var user User
	user = GetUserByID(session.User_ID)
	return user
}

func GetUserByID(userid int) User {

	var user User
	row := Db.QueryRow(`SELECT u.user_id, u.username, u.password, u.email, STRFTIME('%d.%m.%Y', u.reg_datetime) as reg_datetime, u.role_id, r.role,
	CASE WHEN un.unread IS NULL THEN 0 ELSE un.unread END AS unread
	FROM user u
	LEFT JOIN role r ON r.role_id = u.role_id
	LEFT JOIN (SELECT user_id, count(notification_id) unread FROM notification WHERE read=0 GROUP BY user_id) AS un ON un.user_id=u.user_id
	WHERE u.user_id = $1`, userid)
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Date, &user.RoleID, &user.Role, &user.Unread)
	return user
}

func GetSessionByUUID(uuid string) Session {
	var session Session
	row := Db.QueryRow("SELECT * FROM session WHERE uuid = $1", uuid)
	//Put the values in a struct
	row.Scan(&session.Session_ID, &session.Datetime, &session.User_ID, &session.Uuid)
	return session
}
