package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

type commentRequest struct {
	LastEarliestcomment string `json:"lastEarliestcomment"`
}

func SubmitComment(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var comment data.Comment

	err = json.Unmarshal(body, &comment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	valid, err := comment.IsValid()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if valid {
		err := comment.Insert()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}

type CommentInfo struct {
	Comment *data.Comment `json:"comment"`
	User *data.User `json:"user"`
}

func LatestComments(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var lastEarliestComment = queryParams["lastEarliestComment"][0]

	comments, err := data.LatestComments(lastEarliestComment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	commentsInfo := []CommentInfo{}

	for _, comment := range comments {
		var currentComment CommentInfo
		currentComment.Comment = comment

		// Get User by UserID
		_, user, err := data.GetUser("ID", comment.UserID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.Password = []byte("")
		user.Email = ""
		currentComment.User = user

		commentsInfo = append(commentsInfo, currentComment)
	}

	data, err := json.Marshal(commentsInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
