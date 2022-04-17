package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

func submitPost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var post data.Post

	err = json.Unmarshal(body, &post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	post.Insert()
}
