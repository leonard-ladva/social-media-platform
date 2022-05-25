package chat

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
	"git.01.kood.tech/Rostislav/real-time-forum/dataHelpers"
	"github.com/gorilla/websocket"
)

type WebsocketMessage struct {
	Type    string    `json:"type"` // one of 'auth', 'message', 'userOffline', 'userOnline'
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

func authenticate(conn *websocket.Conn, user data.User) {
	client := new(Client)
	client.Id = ClientID(user.ID)
	client.Conn = conn

	GC.Add(client)
	// defer GC.Del(client.Id)

	fmt.Println(GC.list())
}

func handleMessage(conn *websocket.Conn, wsMsg WebsocketMessage, messageType int) error {
	messageTime := data.CurrentTime()

	valid, err := wsMsg.valid()
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("chat.handleMessage: websocket message not valid")
	}

	err = wsMsg.prepareChat(messageTime)
	if err != nil {
		return err
	}

	err = wsMsg.prepareMessage(messageTime)
	if err != nil {
		return err
	}

	err = wsMsg.sendToClient(messageType)
	if err != nil {
		return err
	}
	return nil
}

func (msg WebsocketMessage) valid() (bool, error) {
	senderExsists, _, err := data.GetUser("ID", msg.User.ID)
	if err != nil {
		return false, err
	}

	receiverExsists, _, err := data.GetUser("ID", msg.To)
	if err != nil {
		return false, err
	}

	if !senderExsists || !receiverExsists || msg.Message == "" {
		return false, nil
	}
	return true, nil
}

func (msg WebsocketMessage) prepareChat(msgTime int64) error {
	var chat data.Chat
	chatID, err := dataHelpers.ChatID(msg.User.ID, msg.To)
	if err != nil {
		return err
	}
	chat.ID = chatID
	chat.LastMessageTime = msgTime
	fmt.Println("before: ", chat)
	chatExists, err := chat.Exists()
	if err != nil {
		return err
	}
	fmt.Println("after:", chat)

	if chatExists {
		// chat, err := chat.Get()
		// if err != nil {
		// 	return err
		// }
		// chat.LastMessageTime = msgTime

		err = chat.Update()
		if err != nil {
			return err
		}
	} else {
		err = chat.Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

func (msg WebsocketMessage) prepareMessage(msgTime int64) error {
	var message data.Message
	chatID, err := dataHelpers.ChatID(msg.User.ID, msg.To)
	if err != nil {
		return err
	}
	message.ChatID = chatID
	message.UserID = msg.User.ID
	message.Content = msg.Message
	message.CreatedAt = msgTime

	err = message.Insert()
	if err != nil {
		return err
	}
	return nil
}

func (msg WebsocketMessage) sendToClient(messageType int) error {
	_, online := GC.data[ClientID(msg.To)]
	if online {
		err := GC.data[ClientID(msg.To)].Conn.WriteMessage(messageType, []byte(msg.Message))
		if err != nil {
			return err
		}
	}
	return nil
}
