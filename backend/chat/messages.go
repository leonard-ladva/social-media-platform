package chat

import (
	"encoding/json"
	"errors"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
	"github.com/gorilla/websocket"
)

type WebsocketMessage struct {
	Type    string       `json:"type"` // one of 'auth', 'message', 'offline', 'online'
	UserID  string       `json:"userId"`
	Message data.Message `json:"message"`
}

func handleMessage(conn *websocket.Conn, wsMsg WebsocketMessage, messageType int) error {
	message := wsMsg.Message
	message.CreatedAt = data.CurrentTime()

	chatID, err := data.ChatID(message.UserID, message.ReceiverID)
	if err != nil {
		return err
	}
	message.ChatID = chatID

	valid, err := message.Valid()
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("chat.handleMessage: websocket message not valid")
	}

	err = chatToDB(message)
	if err != nil {
		return err
	}

	err = message.Insert()
	if err != nil {
		return err
	}

	wsMsg.Message = message
	err = wsMsg.sendToClient(messageType)
	if err != nil {
		return err
	}
	return nil
}

func chatToDB(msg data.Message) error {
	var chat data.Chat
	chat.ID = msg.ChatID
	chat.LastMessageTime = msg.CreatedAt

	chatExists, err := chat.Exists()
	if err != nil {
		return err
	}
	if chatExists {
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

func (wsMsg WebsocketMessage) sendToClient(messageType int) error {
	jsonMsg, err := json.Marshal(wsMsg)
	if err != nil {
		return err
	}

	// Send message to receiver
	_, online := GC.Data[ClientID(wsMsg.Message.ReceiverID)]
	if online {
		err := GC.Data[ClientID(wsMsg.Message.ReceiverID)].Conn.WriteMessage(messageType, jsonMsg)
		if err != nil {
			return err
		}
	}
	// If sender is also receiver return to not send message to user twice
	if wsMsg.Message.UserID == wsMsg.Message.ReceiverID {
		return nil
	}

	// Send message back to sender to display
	err = GC.Data[ClientID(wsMsg.Message.UserID)].Conn.WriteMessage(messageType, jsonMsg)
	if err != nil {
		return err
	}
	return nil
}

func (wsMsg WebsocketMessage) broadcast(messageType int) error {
	jsonMsg, err := json.Marshal(wsMsg)
	if err != nil {
		return err
	}

	for _, client := range GC.list() {
		client.Conn.WriteMessage(messageType, jsonMsg)
	}
	return nil
}
