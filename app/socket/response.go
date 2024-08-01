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

type Response[T any] struct {
	Event SocketEvent `json:"event"`
	Data  T           `json:"data,omitempty"`
}

func NewSocketResponse[T any](event SocketEvent, data T) *Response[T] {
	return &Response[T]{
		Event: event,
		Data:  data,
	}
}
