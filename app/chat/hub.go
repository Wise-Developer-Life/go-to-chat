// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	"github.com/ugurcsen/gods-generic/lists/arraylist"
)

type HubInterface interface {
	RegisterClient(username string, client *Client)
	GetClientFromUsername(username string) *Client
	IsUserConnected(username string) bool
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	//retrieve client by username
	userToClient map[string]*Client

	rooms *arraylist.List[*Room]

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func (hub *Hub) GetClientFromUsername(username string) *Client {
	if client, ok := hub.userToClient[username]; ok {
		return client
	}
	return nil
}

func (hub *Hub) IsUserConnected(username string) bool {
	return hub.GetClientFromUsername(username) != nil
}

func (hub *Hub) RegisterClient(username string, client *Client) {
	hub.userToClient[username] = client
	hub.register <- client
}

func newHub() *Hub {
	return &Hub{
		broadcast:    make(chan []byte),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
		userToClient: make(map[string]*Client),
		rooms:        arraylist.New[*Room](),
	}
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true

		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}
		}
	}
}
