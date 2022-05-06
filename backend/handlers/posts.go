package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

type postRequest struct {
	LastEarliestPost string `json:"lastEarliestPost"`
}

func Submit(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var post data.Post

	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(post)
	err = post.Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type PostInfo struct {
	Post *data.Post `json:"post"`
	User *data.User `json:"user"`
	Tag  *data.Tag  `json:"tag"`
}

func LatestPosts(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	queryParams := r.URL.Query()
	var lastEarliestPost = queryParams["lastEarliestPost"][0]

	posts, err := data.LatestPosts(lastEarliestPost)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	postsInfo := []PostInfo{}

	for _, post := range posts {
		var currentPost PostInfo
		currentPost.Post = post
		// Get tag by TagID
		tag, err := data.GetTagByID(post.TagID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		currentPost.Tag = tag
		// Get User by UserID
		_, user, err := data.GetUser("ID", post.UserID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.Password = []byte("")
		user.Email = ""
		currentPost.User = user

		postsInfo = append(postsInfo, currentPost)
	}

	data, err := json.Marshal(postsInfo)
	// data, err := json.Marshal(posts)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
