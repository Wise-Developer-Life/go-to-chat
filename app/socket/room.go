package socket

import (
	"github.com/ugurcsen/gods-generic/sets"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

type Room interface {
	GetID() string
	GetSize() int
	IsEmpty() bool
	Add(client Client, otherClients ...Client)
	Remove(client Client, otherClients ...Client)
	Broadcast(payload any)
}

type RoomImpl struct {
	roomID      string
	clients     sets.Set[Client]
	msgToBeSent chan any
}

func NewRoom(roomID string) Room {
	newRoom := &RoomImpl{
		roomID:      roomID,
		clients:     hashset.New[Client](),
		msgToBeSent: make(chan any),
	}
	go newRoom.run()
	return newRoom
}

func (r *RoomImpl) GetID() string {
	return r.roomID
}

func (r *RoomImpl) GetSize() int {
	return r.clients.Size()
}

func (r *RoomImpl) IsEmpty() bool {
	return r.clients.Empty()
}

func (r *RoomImpl) Add(client Client, otherClients ...Client) {
	clients := append([]Client{client}, otherClients...)
	for _, c := range clients {
		r.clients.Add(c)
		c.AddRoom(r)
	}
}

func (r *RoomImpl) Remove(client Client, otherClients ...Client) {
	clients := append([]Client{client}, otherClients...)
	for _, c := range clients {
		r.clients.Remove(c)
		c.RemoveRoom(r)
	}

	if r.IsEmpty() {
		close(r.msgToBeSent)
	}
}

func (r *RoomImpl) Broadcast(payload any) {
	r.msgToBeSent <- payload
}

func (r *RoomImpl) run() {
	for {
		select {
		case message := <-r.msgToBeSent:
			for _, client := range r.clients.Values() {
				client.Send(message)
			}
		}
	}
}
