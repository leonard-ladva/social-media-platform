package handlers

import (
	"encoding/json"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

func GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	tags, err := data.GetAllTags()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(tags)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
