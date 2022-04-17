package handlers

import "net/http"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-type")
	(*w).Header().Add("Access-Control-Allow-Headers", "Authorization")
}
