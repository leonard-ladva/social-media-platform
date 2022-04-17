package main

import (
	"fmt"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
	"git.01.kood.tech/Rostislav/real-time-forum/handlers"
)

func setupRoutes() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/isUnique", handlers.IsUnique)
	http.HandleFunc("/user", handlers.Session)
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
