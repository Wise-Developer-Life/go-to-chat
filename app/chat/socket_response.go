package chat

import "go-to-chat/app/user"

type SocketEvent string

const (
	SocketEventMessage SocketEvent = "message"
	SocketEventMatched SocketEvent = "matched"
	SocketEventJoin    SocketEvent = "join"
	SocketEventLeave   SocketEvent = "leave"
	SocketEventError   SocketEvent = "error"
)

type SocketResponse[T any] struct {
	Event SocketEvent `json:"event"`
	Data  T           `json:"data,omitempty"`
}

type SocketMatchUserResponse struct {
	MatchedUser *user.UserResponse `json:"matched_user"`
}
