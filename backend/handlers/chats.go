package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

func LatestMessages(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var lastEarliest = queryParams["lastEarliest"][0]
	var chatID = queryParams["chatID"][0]

	chat := &data.Chat{}
	chat.ID = chatID

	exists, err := chat.Exists()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exists {
		log.Println("Chat does not exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	messages, err := data.GetLatestMessages(lastEarliest, chatID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(messages)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
