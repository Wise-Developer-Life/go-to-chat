package chat

import (
	"encoding/json"
)

type RoomInterface interface {
	Join(...*Client)
	Leave(...*Client)
	Send(payload any)
}

type Room struct {
	clients map[*Client]bool

	broadcast chan []byte

	register chan *Client

	unregister chan *Client
}

func (room *Room) Join(client ...*Client) {
	for _, c := range client {
		if c != nil {
			room.register <- c
		}
	}
}

func (room *Room) Leave(client ...*Client) {
	for _, c := range client {
		room.unregister <- c
	}
}

func (room *Room) Send(payload any) {
	payloadBytes, _ := json.Marshal(payload)
	room.broadcast <- payloadBytes
}

func (room *Room) run() {
	for {
		select {
		case client := <-room.register:
			room.clients[client] = true
		case client := <-room.unregister:
			if _, ok := room.clients[client]; ok {
				delete(room.clients, client)
			}
		case message := <-room.broadcast:
			for client := range room.clients {
				select {
				case client.send <- message:
				default:
					delete(room.clients, client)
				}
			}
		}
	}
}

func CreateRoom() *Room {
	room := &Room{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}

	hub.rooms.Add(room)
	go room.run()

	return room
}

func (room *Room) getSize() int {
	return len(room.clients)
}

func (room *Room) isFull() bool {
	return room.getSize() >= 2
}
