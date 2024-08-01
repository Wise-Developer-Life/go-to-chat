package socket

import (
	"fmt"
	"github.com/ugurcsen/gods-generic/sets"
	"github.com/ugurcsen/gods-generic/sets/hashset"
	"log"
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

		msg := fmt.Sprintf("Client %s join room %s", c.GetID(), r.GetID())
		log.Println(msg)
		r.Broadcast(NewSocketResponse(SocketEventJoined, msg))
	}
}

func (r *RoomImpl) Remove(client Client, otherClients ...Client) {
	clients := append([]Client{client}, otherClients...)
	for _, c := range clients {
		r.clients.Remove(c)
		c.RemoveRoom(r)

		msg := fmt.Sprintf("Client %s removed from room %s", c.GetID(), r.GetID())
		log.Println(msg)
		r.Broadcast(NewSocketResponse(SocketEventLeft, msg))
	}

	if r.IsEmpty() {
		fmt.Sprintf("Room %s is empty. Closing the room", r.GetID())
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
