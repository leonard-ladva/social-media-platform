package main

import (
	"fmt"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/database"
)

func setupRoutes() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	//http.HandleFunc("/", homeHandler)
}

func main() {
	database.Connect()
	fmt.Println("Database Connected")

	setupRoutes()
	http.ListenAndServe(":9100", nil)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-type")
}
