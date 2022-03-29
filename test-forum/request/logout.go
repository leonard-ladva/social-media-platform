package request

import (
	"forum-test/database"
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request, user database.User) {

	database.Db.Exec("DELETE FROM sessions WHERE user_id = $1", user.ID)
	log.Println("Deleting session: ", user.Username)
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
