package socket

import (
	"github.com/ugurcsen/gods-generic/maps"
	"github.com/ugurcsen/gods-generic/maps/hashmap"
	"sync"
)

type Hub interface {
	Register(client Client)
	Unregister(client Client)
	GetClient(clientID string) Client
	IsClientConnected(clientID string) bool
	Broadcast(payload any)
}

type HubImpl struct {
	clientsMap   maps.Map[string, Client]
	broadcastMsg chan any
}

var (
	instance *HubImpl
	once     sync.Once
)

func GetHubInstance() Hub {
	once.Do(func() {
		instance = &HubImpl{
			clientsMap:   hashmap.New[string, Client](),
			broadcastMsg: make(chan any),
		}
		go instance.run()
	})
	return instance
}

func (hub *HubImpl) Register(client Client) {
	hub.clientsMap.Put(client.GetID(), client)
}

func (hub *HubImpl) Unregister(client Client) {
	hub.clientsMap.Remove(client.GetID())
}

func (hub *HubImpl) GetClient(clientID string) Client {
	if client, found := hub.clientsMap.Get(clientID); found {
		return client
	}
	return nil
}

func (hub *HubImpl) IsClientConnected(clientID string) bool {
	return hub.GetClient(clientID) != nil
}

func (hub *HubImpl) Broadcast(payload any) {
	hub.broadcastMsg <- payload
}

func (hub *HubImpl) run() {
	for {
		select {
		case message := <-hub.broadcastMsg:
			for _, client := range hub.clientsMap.Values() {
				client.Send(message)
			}
		}
	}
}
