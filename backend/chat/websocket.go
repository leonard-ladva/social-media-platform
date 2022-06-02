package chat

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
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
	var clientID ClientID
	for {
		fmt.Println("\n------\n", GC.list(), "\n------")
		messageType, body, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			WebsocketClosed(clientID, conn)
			return
		}
		var wsMsg WebsocketMessage
		json.Unmarshal(body, &wsMsg)

		switch wsMsg.Type {
		case "auth":
			clientID, err = authenticate(conn, wsMsg.UserID)
			if err != nil {
				log.Println(err)
				return
			}
			WebsocketOpened(clientID)
		case "message":
			err := handleMessage(conn, wsMsg, messageType)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func WebsocketClosed(cid ClientID, conn *websocket.Conn) {
	GC.Del(cid)
	conn.Close()

	var wsMsg WebsocketMessage
	wsMsg.Type = "offline"
	wsMsg.UserID = string(cid)

	wsMsg.broadcast(1)
}

func WebsocketOpened(cid ClientID) {
	var wsMsg WebsocketMessage
	wsMsg.Type = "online"
	wsMsg.UserID = string(cid)

	wsMsg.broadcast(1)
}

func authenticate(conn *websocket.Conn, UserID string) (ClientID, error) {
	client := new(Client)
	client.Id = ClientID(UserID)
	client.Conn = conn
	exsists, user, err := data.GetUser("ID", UserID)
	if err != nil {
		return client.Id, err
	}
	if !exsists {
		return client.Id, errors.New("User Doesn't exist")
	}
	client.Nickname = user.Nickname

	GC.Add(client)

	return client.Id, nil
}
