package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var GC = globalClients{data: make(map[ClientID]*Client), RWMutex: &sync.RWMutex{}}

type globalClients struct {
	data map[ClientID]*Client
	*sync.RWMutex
}

type Client struct {
	Conn *websocket.Conn
	Name string
	Id   ClientID
}

type ClientID string

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

type ClientList struct {
	Name string   `json:"client_name"`
	ID   ClientID `json:"client_id"`
}

func (gc *globalClients) List() []ClientList {

	gc.RLock()
	defer gc.RUnlock()

	out := []ClientList{}

	for _, cl := range gc.data {
		out = append(out, ClientList{Name: cl.Name, ID: cl.Id})
	}

	return out
}

func ListAllClients(rw http.ResponseWriter, r *http.Request) {

	data := GC.List()

	enc, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error occured while listing clients : %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Some error occured"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(enc)
}
