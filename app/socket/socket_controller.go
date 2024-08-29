package socket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-to-chat/app/chat"
	"go-to-chat/app/match/response"
	user2 "go-to-chat/app/user"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:3000"
	},
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := r.URL.Query().Get("user")

	client := NewClient(user, conn)
	hub := GetHubInstance()
	hub.Register(client)
	client.Send(NewSocketMessage[any](SocketEventConnected, nil))
	log.Println(fmt.Sprintf("Client %s connected", client.GetID()))

	chatRoomService := chat.GetChatRoomServiceInstance()
	roomOfUser, err := chatRoomService.GetRoomOfUser(user)

	if err != nil {
		log.Println("Error getting chat room: ", err)
		return
	}

	// FIXME: refactor here in future, use API instead of direct call
	usersInRoom := roomOfUser.GetUsers()
	var receiver string
	for _, userInRoom := range usersInRoom {
		if userInRoom != user {
			receiver = userInRoom
			break
		}
	}

	userService := user2.NewUserService(user2.NewUserRepository())
	matchedUser, err := userService.GetUserByEmail(receiver)

	if err != nil {
		log.Println("Error getting user: ", err)
		return
	}

	matchedSocketMessage := NewSocketMessage(SocketEventMatched, &response.MatchUserResponse{
		MatchedUser: user2.NewUserResponse(matchedUser),
	})

	client.Send(matchedSocketMessage)

	messages, err := chatRoomService.GetMessages(roomOfUser.ID)
	socketsMessages := make([]*SocketMessage[*ChatMessage], 0)

	for _, message := range messages {
		socketsMessages = append(socketsMessages, NewSocketMessage[*ChatMessage](SocketEventMessage, &ChatMessage{
			Sender:   message.Sender,
			Message:  message.Message,
			Receiver: receiver,
		}))
	}

	for _, socketMessage := range socketsMessages {
		client.Send(socketMessage)
	}
}

func HandleWebSocket() gin.HandlerFunc {
	return func(c *gin.Context) {
		serveWs(c.Writer, c.Request)
	}
}
