package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
)

type postRequest struct {
	LastEarliestPost string `json:"lastEarliestPost"`
}

func SubmitPost(w http.ResponseWriter, r *http.Request) {
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

	valid, err := post.IsValid()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	if valid {
		err := post.Insert()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}

type PostInfo struct {
	Post *data.Post `json:"post"`
	User *data.User `json:"user"`
	Tag  *data.Tag  `json:"tag"`
}

func LatestPosts(w http.ResponseWriter, r *http.Request) {
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

func GetPost(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var postID = queryParams["postId"][0]

	post, err := data.GetPost(postID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	postInfo := PostInfo{}

	postInfo.Post = post
	// Get tag by TagID
	tag, err := data.GetTagByID(post.TagID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	postInfo.Tag = tag
	// Get User by UserID
	_, user, err := data.GetUser("ID", post.UserID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Password = []byte("")
	user.Email = ""
	postInfo.User = user

	data, err := json.Marshal(postInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

