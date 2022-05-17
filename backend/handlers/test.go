package handlers

import (
	"fmt"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/chat"
)

func Test(cl *chat.Client, w http.ResponseWriter, r *http.Request) {
	fmt.Println(cl)
}
