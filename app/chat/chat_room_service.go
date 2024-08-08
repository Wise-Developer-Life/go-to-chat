package chat

import (
	"encoding/json"
	"github.com/ugurcsen/gods-generic/maps"
	"github.com/ugurcsen/gods-generic/maps/hashmap"
	"go-to-chat/app/model"
	"go-to-chat/app/socket"
	"go-to-chat/app/utility"
	"log"
)

type ChatRoomService interface {
	CreateChatRoom(chatRoomName string) (*model.ChatRoom, error)
	JoinChatRoom(chatRoomName string, user string) error
	LeaveChatRoom(chatRoomName string, user string) error
	SendMessage(chatRoomName string, user string, message string) error
}

type chatRoomServiceImpl struct {
	roomsMap    maps.Map[string, *model.ChatRoom]
	redisClient utility.RedisClient
}

func newChatRoomService() ChatRoomService {
	redisClient, err := utility.GetRedisClient(utility.TypeDefault)

	if err != nil {
		log.Println("Error getting redis client's instance: ", err)
		return nil
	}

	return &chatRoomServiceImpl{
		roomsMap:    hashmap.New[string, *model.ChatRoom](),
		redisClient: redisClient,
	}
}

var serviceInstance ChatRoomService

func GetChatRoomServiceInstance() ChatRoomService {
	if serviceInstance == nil {
		serviceInstance = newChatRoomService()

		if serviceInstance == nil {
			log.Println("Error getting chat room service instance")
		}
	}
	return serviceInstance
}

func (service *chatRoomServiceImpl) CreateChatRoom(chatRoomName string) (*model.ChatRoom, error) {
	chatRoom := model.NewChatRoom(chatRoomName)
	_ = service.redisClient.Set(chatRoomName, chatRoom, 0)
	return chatRoom, nil
}

func (service *chatRoomServiceImpl) JoinChatRoom(chatRoomName string, user string) error {
	chatRoomStr, err := service.redisClient.Get(chatRoomName)

	if err != nil {
		log.Println("Error getting chat room: ", err)
		return err
	}

	var chatRoom *model.ChatRoom
	err = json.Unmarshal([]byte(chatRoomStr), &chatRoom)

	if err != nil {
		log.Println("Error parsing chat room: ", err)
		return err
	}

	chatRoom.Users.Add(user)
	_ = service.redisClient.Set(chatRoomName, chatRoom, 0)

	return nil
}

func (service *chatRoomServiceImpl) LeaveChatRoom(chatRoomName string, user string) error {
	chatRoomStr, err := service.redisClient.Get(chatRoomName)

	if err != nil {
		log.Println("Error getting chat room: ", err)
		return err
	}

	var chatRoom *model.ChatRoom
	err = json.Unmarshal([]byte(chatRoomStr), &chatRoom)

	if err != nil {
		log.Println("Error parsing chat room: ", err)
		return err
	}

	chatRoom.Users.Remove(user)
	_ = service.redisClient.Set(chatRoomName, chatRoom, 0)

	return nil
}

func (service *chatRoomServiceImpl) SendMessage(chatRoomName string, user string, message string) error {
	chatRoomStr, err := service.redisClient.Get(chatRoomName)

	if err != nil {
		log.Println("Error getting chat room: ", err)
		return err
	}

	var chatRoom *model.ChatRoom
	err = json.Unmarshal([]byte(chatRoomStr), &chatRoom)

	if err != nil {
		log.Println("Error parsing chat room: ", err)
		return err
	}

	if !chatRoom.Users.Contains(user) {
		log.Println("User is not in the chat room")
		return nil
	}

	chatMessage := model.NewChatMessage(user, message)
	socketHub := socket.GetHubInstance()
	for _, userInRoom := range chatRoom.Users.Values() {
		socketClient := socketHub.GetClient(userInRoom)
		if socketClient != nil {
			socketClient.Send(chatMessage)
		}
	}

	err = service.redisClient.ZAdd(chatRoomName, chatMessage, float64(chatMessage.Timestamp))

	return err
}
