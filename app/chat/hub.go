// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	"github.com/ugurcsen/gods-generic/lists/arraylist"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	rooms *arraylist.List[*Room]

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rooms:      arraylist.New[*Room](),
	}
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true

			lastRoom, _ := hub.rooms.Get(hub.rooms.Size() - 1)
			if hub.rooms.Empty() || lastRoom.isFull() {
				room := initializeRoom()
				room.register <- client
				hub.rooms.Add(room)
			} else {
				lastRoom.register <- client
			}

		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				rooms := hub.rooms

				rooms.Each(func(index int, room *Room) {
					room.unregister <- client
				})

				delete(hub.clients, client)
				close(client.send)
			}
			//case message := <-hub.broadcast:
			//for client := range hub.clients {
			//	select {
			//	case client.send <- message:
			//	default:
			//		close(client.send)
			//		delete(hub.clients, client)
			//	}
			//}
		}
	}
}
