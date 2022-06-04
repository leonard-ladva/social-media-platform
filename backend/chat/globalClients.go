package chat

import (
	"sync"

	"github.com/gorilla/websocket"
)

var GC = globalClients{Data: make(map[ClientID]*Client), RWMutex: &sync.RWMutex{}}

type globalClients struct {
	Data map[ClientID]*Client
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
	gc.Data[cl.Id] = cl
}

func (gc *globalClients) Del(cid ClientID) {
	gc.Lock()
	defer gc.Unlock()
	delete(gc.Data, cid)
}

func (gc *globalClients) list() []Client {

	gc.RLock()
	defer gc.RUnlock()

	out := []Client{}

	for _, cl := range gc.Data {
		out = append(out, Client{Id: cl.Id, Nickname: cl.Nickname, Conn: cl.Conn})
		// out = append(out, Client{Nickname: cl.Nickname})
	}

	return out
}
