package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

func EnableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Headers", "Content-type")
		(w).Header().Add("Access-Control-Allow-Headers", "Authorization")

		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// r.Header.Get("Authorization") returns "Bearer <ActualToken>", so we only need the second part
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]
		fmt.Println("token: ", token)

		session, err := data.GetSession(token)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := &data.User{}
		_, user, err = data.GetUser("ID", session.UserID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data.CurrentUser = user
		next(w, r)
	}
}
