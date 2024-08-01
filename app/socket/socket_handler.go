package socket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

	log.Println(fmt.Sprintf("Client %s connected", client.GetID()))
}

func HandleWebSocket() gin.HandlerFunc {
	return func(c *gin.Context) {
		serveWs(c.Writer, c.Request)
	}
}
