package model

import "time"

type ChatMessage struct {
	Sender    string `json:"user"`
	Message   string `json:"message"`
	Timestamp int64  `json:"time"`
}

func NewChatMessage(sender string, message string) *ChatMessage {
	return &ChatMessage{
		Sender:    sender,
		Message:   message,
		Timestamp: time.Now().Unix(),
	}
}
