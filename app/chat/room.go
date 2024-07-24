package chat

type Room struct {
	clients map[*Client]bool

	broadcast chan []byte

	register chan *Client

	unregister chan *Client
}

func initializeRoom() *Room {
	room := &Room{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}

	go room.run()

	return room
}

func (room *Room) run() {
	for {
		select {
		case client := <-room.register:
			room.clients[client] = true
		case client := <-room.unregister:
			if _, ok := room.clients[client]; ok {
				delete(room.clients, client)
			}
		case message := <-room.broadcast:
			for client := range room.clients {
				select {
				case client.send <- message:
				default:
					delete(room.clients, client)
				}
			}
		}
	}
}

func (room *Room) getSize() int {
	return len(room.clients)
}

func (room *Room) isFull() bool {
	return room.getSize() >= 2
}
