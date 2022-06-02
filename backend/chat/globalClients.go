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
	Nickname string
	Id ClientID
}

type ClientID string

type ClientList struct {
	ID   ClientID
	Conn *websocket.Conn
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

func (gc *globalClients) list() []Client {

	gc.RLock()
	defer gc.RUnlock()

	out := []Client{}

	for _, cl := range gc.data {
		out = append(out, Client{Id: cl.Id, Nickname: cl.Nickname, Conn: cl.Conn})
		// out = append(out, Client{Nickname: cl.Nickname})
	}

	return out
}
