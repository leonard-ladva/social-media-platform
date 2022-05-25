package chat

import (
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
	// Name string
	Id ClientID
}

type ClientID string

type ClientList struct {
	Name string   `json:"client_name"`
	ID   ClientID `json:"client_id"`
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

func (gc *globalClients) list() []ClientList {

	gc.RLock()
	defer gc.RUnlock()

	out := []ClientList{}

	for _, cl := range gc.data {
		out = append(out, ClientList{ID: cl.Id})
	}

	return out
}
