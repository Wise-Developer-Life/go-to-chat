package model

import (
	"encoding/json"
	"errors"
)

type ChatRoom struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Users    map[string]bool `json:"users"`
	Messages []*ChatMessage  `json:"messages"`
}

func NewChatRoom(name string) *ChatRoom {
	return &ChatRoom{
		ID:    name,
		Name:  name,
		Users: make(map[string]bool),
	}
}

func (c *ChatRoom) AddUser(user string) {
	c.Users[user] = true
}

func (c *ChatRoom) RemoveUser(user string) {
	c.Users[user] = false
}

func (c *ChatRoom) AddMessage(message *ChatMessage) error {
	if !c.Users[message.Sender] {
		return errors.New("user not in chat room")
	}
	c.Messages = append(c.Messages, message)
	return nil
}

func (c *ChatRoom) GetMessages() []*ChatMessage {
	return c.Messages
}

func (c *ChatRoom) ContainsUser(user string) bool {
	return c.Users[user]
}

func (c *ChatRoom) GetUsers() []string {
	users := make([]string, 0)
	for user, _ := range c.Users {
		users = append(users, user)
	}
	return users
}

func (c *ChatRoom) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *ChatRoom) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
