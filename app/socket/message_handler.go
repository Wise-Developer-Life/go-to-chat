package socket

import (
	"encoding/json"
	"fmt"
	"go-to-chat/app/chat"
	"go-to-chat/app/model"
	"log"
)

func HandleSocketInputMessage(client Client, event SocketEvent, data any) {
	dataBytes, _ := json.Marshal(data)

	if event == SocketEventMessage {
		log.Println(fmt.Sprintf("Message received from %s: %s", client.GetID(), data))
		var chatMessage ChatMessage
		err := json.Unmarshal(dataBytes, &chatMessage)

		if err != nil {
			log.Println("Error unmarshalling chat message: ", err)
			return
		}

		handleChatMessage(client, &chatMessage)
	}
}

func handleChatMessage(client Client, chatMessage *ChatMessage) {
	chatRoomService := chat.GetChatRoomServiceInstance()
	hub := GetHubInstance()

	chatRoom, err := chatRoomService.GetRoomOfUser(chatMessage.Sender)

	if err != nil {
		log.Println("Error getting chat room: ", err)
		return
	}
	log.Println(fmt.Sprintf("Chat room: %s", chatRoom.ID))

	for _, userInRoom := range chatRoom.GetUsers() {
		socketClient := hub.GetClient(userInRoom)
		if socketClient != nil {
			socketClient.Send(NewSocketMessage(SocketEventMessage, chatMessage))
		}
	}

	_ = chatRoomService.SaveMessage(chatRoom.ID, model.NewChatMessage(chatMessage.Sender, chatMessage.Message))
}
