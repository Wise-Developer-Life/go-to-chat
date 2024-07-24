package chat

import (
	"github.com/gin-gonic/gin"
	"log"
)

var hub *Hub = nil

func HandleChatSocket() gin.HandlerFunc {
	if hub == nil {
		log.Println("initialize hub")
		hub = newHub()
		go hub.run()
	}
	return func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	}
}
