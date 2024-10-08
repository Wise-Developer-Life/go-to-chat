package socket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client interface {
	GetID() string
	Send(payload any)
	Close() error
}

type ClientImpl struct {
	connection  *websocket.Conn
	clientID    string
	msgToBeSent chan any
}

func NewClient(clientID string, connection *websocket.Conn) Client {
	newClient := &ClientImpl{
		clientID:    clientID,
		connection:  connection,
		msgToBeSent: make(chan any),
	}

	connection.SetCloseHandler(func(code int, text string) error {
		log.Println(fmt.Sprintf("Client %s connection closed. Code: %d, Text: %s", newClient.GetID(), code, text))
		newClient.Close()
		hub := GetHubInstance()
		hub.Unregister(newClient)
		return nil
	})

	go newClient.writePump()
	go newClient.readPump()

	return newClient
}

func (client *ClientImpl) GetID() string {
	return client.clientID
}

func (client *ClientImpl) Send(payload any) {
	client.msgToBeSent <- payload
}

func (client *ClientImpl) Close() error {
	return nil
}

func (client *ClientImpl) readPump() {
	defer func() {
		client.connection.Close()
	}()

	client.connection.SetReadLimit(maxMessageSize)
	client.connection.SetReadDeadline(time.Now().Add(pongWait))

	client.connection.SetPongHandler(func(string) error {
		client.connection.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		var message SocketMessage[any]
		err := client.connection.ReadJSON(&message)

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		HandleSocketInputMessage(client, message.Event, message.Data)
	}
}

func (client *ClientImpl) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.connection.Close()
	}()
	for {
		select {
		case message, ok := <-client.msgToBeSent:
			client.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				client.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			messageBytes, _ := json.Marshal(message)
			log.Println(fmt.Sprintf("Message sent to %s: %s", client.GetID(), string(messageBytes)))
			err := client.connection.WriteJSON(message)
			if err != nil {
				return
			}

		case <-ticker.C:
			client.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
