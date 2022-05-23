package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
	"github.com/gorilla/websocket"
)

var GC = globalClients{data: make(map[ClientID]*Client), RWMutex: &sync.RWMutex{}}

type globalClients struct {
	data map[ClientID]*Client
	*sync.RWMutex
}

type Client struct {
	Conn *websocket.Conn
	// Name string
	Id ClientID
}

type ClientID string

type ClientList struct {
	Name string   `json:"client_name"`
	ID   ClientID `json:"client_id"`
}
type wsMsg struct {
	Type    string    `json:"type"`
	User    data.User `json:"user"`
	Message string    `json:"message"`
	To      string    `json:"to"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := data.GetAllUsers()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If a user exists in the WebSocket Global Clients map then mark as active
	for _, user := range users {
		_, ok := GC.data[ClientID(user.ID)]
		if ok {
			user.Active = true
		}
	}

	data, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (gc *globalClients) Add(cl *Client) {
	gc.Lock()
	defer gc.Unlock()
	gc.data[cl.Id] = cl
}

func (gc *globalClients) Del(cid ClientID) {
	gc.Lock()
	defer gc.Unlock()
	delete(gc.data, cid)
}

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
		var msgBody wsMsg
		json.Unmarshal(body, &msgBody)

		switch msgBody.Type {
		case "auth":
			authenticate(conn, msgBody.User)
			break
		case "message":
			handleMessage(conn, msgBody, messageType)
		}
		fmt.Println(string(body))

		if err := conn.WriteMessage(messageType, body); err != nil {
			log.Println(err)
			return
		}

	}
}

func authenticate(conn *websocket.Conn, user data.User) {
	fmt.Println(user)

	client := new(Client)
	client.Id = ClientID(user.ID)
	client.Conn = conn

	GC.Add(client)
	// defer GC.Del(client.Id)

	fmt.Println(GC.list())

}

func handleMessage(conn *websocket.Conn, msgBody wsMsg, messageType int) {
	_, online := GC.data[ClientID(msgBody.To)]
	if online {
		err := GC.data[ClientID(msgBody.To)].Conn.WriteMessage(messageType, []byte(msgBody.Message))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (gc *globalClients) list() []ClientList {

	gc.RLock()
	defer gc.RUnlock()

	out := []ClientList{}

	for _, cl := range gc.data {
		out = append(out, ClientList{ID: cl.Id})
	}

	return out
}
