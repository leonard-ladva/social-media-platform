package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

func GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := data.GetAllTags()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(tags)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
