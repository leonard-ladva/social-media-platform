package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Client Successfully Connected...")
	wsReader(ws)
}

func wsReader(conn *websocket.Conn) {
	for {
		messageType, body, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var wsMsg WebsocketMessage
		json.Unmarshal(body, &wsMsg)

		switch wsMsg.Type {
		case "auth":
			authenticate(conn, wsMsg.UserID)
		case "message":
			err := handleMessage(conn, wsMsg, messageType)
			if err != nil {
				log.Println(err)
				return
			}
		}
		fmt.Println(wsMsg)

		// if err := conn.WriteMessage(messageType, body); err != nil {
		// 	log.Println(err)
		// 	return
		// }

	}
}
