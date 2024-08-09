package socket

type SocketEvent string

const (
	SocketEventMessage   SocketEvent = "message"
	SocketEventMatched   SocketEvent = "matched"
	SocketEventConnected SocketEvent = "connected"
	SocketEventError     SocketEvent = "error"
	SocketEventJoined    SocketEvent = "joined"
	SocketEventLeft      SocketEvent = "left"
)

type SocketMessage[T any] struct {
	Event SocketEvent `json:"event"`
	Data  T           `json:"data,omitempty"`
}

func NewSocketMessage[T any](event SocketEvent, data T) *SocketMessage[T] {
	return &SocketMessage[T]{
		Event: event,
		Data:  data,
	}
}

type ChatMessage struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	//RoomID   string `json:"room_id"`
	Message string `json:"text"`
}
