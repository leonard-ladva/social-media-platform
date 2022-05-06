package main

import (
	"fmt"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
	"git.01.kood.tech/Rostislav/real-time-forum/handlers"
)

func setupRoutes() {
	// Authentication
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/isUnique", handlers.IsUnique)
	http.HandleFunc("/user", handlers.Session)
	// Posts
	http.HandleFunc("/submitPost", handlers.Submit)
	http.HandleFunc("/latestPosts", handlers.LatestPosts)
	// Tags
	http.HandleFunc("/tags", handlers.GetTagsHandler)
	// Chats
	http.HandleFunc("/users", handlers.GetAllUsers)
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
