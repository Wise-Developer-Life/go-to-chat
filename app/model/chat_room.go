package model

import (
	"github.com/ugurcsen/gods-generic/sets"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

type ChatRoom struct {
	ID    string
	Name  string
	Users sets.Set[string]
}

func NewChatRoom(name string) *ChatRoom {
	return &ChatRoom{
		ID:    name,
		Name:  name,
		Users: hashset.New[string](),
	}
}
