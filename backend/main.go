package main

import (
	"fmt"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/chat"
	"git.01.kood.tech/Rostislav/real-time-forum/data"
	"git.01.kood.tech/Rostislav/real-time-forum/handlers"
	mid "git.01.kood.tech/Rostislav/real-time-forum/middleware"
)

var port string

func setupRoutes() {
	// Authentication
	http.HandleFunc("/login", mid.EnableCors(handlers.Login))
	http.HandleFunc("/register", mid.EnableCors(handlers.Register))
	http.HandleFunc("/isUnique", mid.EnableCors(handlers.IsUnique))
	http.HandleFunc("/user", mid.EnableCors(mid.Authenticate(handlers.CurrentUser)))
	// Posts
	http.HandleFunc("/submitPost", mid.EnableCors(mid.Authenticate(handlers.SubmitPost)))
	http.HandleFunc("/latestPosts", mid.EnableCors(mid.Authenticate(handlers.LatestPosts)))
	// Tags
	http.HandleFunc("/tags", mid.EnableCors(mid.Authenticate(handlers.GetTagsHandler)))
	// Chats
	http.HandleFunc("/users", mid.EnableCors(mid.Authenticate(chat.GetAllUsers)))
	http.HandleFunc("/ws", chat.WebSocket)
	http.HandleFunc("/latestMessages", mid.EnableCors(mid.Authenticate(handlers.LatestMessages)))
	// Comments
	http.HandleFunc("/submitComment", mid.EnableCors(mid.Authenticate(handlers.SubmitComment)))
	http.HandleFunc("/latestComments", mid.EnableCors(mid.Authenticate(handlers.LatestComments)))

}

func main() {
	err := data.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")

	setupRoutes()
	http.ListenAndServe(":9100", nil)
}
