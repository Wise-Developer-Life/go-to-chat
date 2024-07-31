package chat

import (
	"github.com/gin-gonic/gin"
	"log"
)

var hub *Hub = nil

func GetHubInstance() *Hub {
	if hub == nil {
		log.Println("initialize hub")
		hub = newHub()
		go hub.run()
	}
	return hub
}

func HandleChatSocket() gin.HandlerFunc {
	hub := GetHubInstance()
	return func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	}
}
