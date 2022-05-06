package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	users, err := data.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
