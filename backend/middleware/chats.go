package middleware

import (
	"fmt"
	"log"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/chat"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

const (
	readBuffSize = 2 << 10
	writeBuffSize
)

const (
	nameHeader = "WS-NAME"
	idHeader   = "WS-ID"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  readBuffSize,
	WriteBufferSize: writeBuffSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ClientMW(next func(cl *chat.Client, rw http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	var cid string
	var conn *websocket.Conn
	var err error

	return func(rw http.ResponseWriter, r *http.Request) {

		qv := r.URL.Query()
		name := qv.Get("name")

		cid = uuid.NewV4().String() //Create a UUID
		h := http.Header{}          //Create a header which we will pass to the client which will contain the UUID
		h.Add(idHeader, cid)

		conn, err = upgrader.Upgrade(rw, r, h) //Upgrade the conenction
		if err != nil {
			log.Printf("Error while upgrading connection: %v", err)
			return
		}

		cl := new(chat.Client) //Create a client pointer
		cl.Conn = conn
		cl.Id = chat.ClientID(cid)
		cl.Name = name

		chat.GC.Add(cl) //Add to global clients struct.
		fmt.Println(chat.GC.List())

		defer chat.GC.Del(cl.Id) //Delete from global clients struct.

		next(cl, rw, r) //Call the passed handler.
	}

}
