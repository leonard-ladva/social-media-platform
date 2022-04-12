package main

import (
	"fmt"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
	"git.01.kood.tech/Rostislav/real-time-forum/handlers"
)

func setupRoutes() {
	// http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
}

func main() {
	data.Connect()
	fmt.Println("Database Connected")

	setupRoutes()
	http.ListenAndServe(":9100", nil)
}
