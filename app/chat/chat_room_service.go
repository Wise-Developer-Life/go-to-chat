package chat

import (
	"go-to-chat/app/model"
	"log"
)

type ChatRoomService interface {
	CreateChatRoom(chatRoomName string) (*model.ChatRoom, error)
	GetChatRoom(chatRoomName string) (*model.ChatRoom, error)
	JoinChatRoom(chatRoomName string, user string) error
	LeaveChatRoom(chatRoomName string, user string) error
}

type chatRoomServiceImpl struct {
	chatRoomRepo RoomRepository
}

func newChatRoomService() ChatRoomService {
	return &chatRoomServiceImpl{
		chatRoomRepo: NewChatRoomRepository(),
	}
}

var serviceInstance ChatRoomService

func GetChatRoomServiceInstance() ChatRoomService {
	if serviceInstance == nil {
		serviceInstance = newChatRoomService()
	}
	return serviceInstance
}

func (service *chatRoomServiceImpl) GetChatRoom(chatRoomName string) (*model.ChatRoom, error) {
	return service.chatRoomRepo.GetChatRoom(chatRoomName)
}

func (service *chatRoomServiceImpl) CreateChatRoom(chatRoomName string) (*model.ChatRoom, error) {
	return service.chatRoomRepo.CreateChatRoom(chatRoomName)
}

func (service *chatRoomServiceImpl) JoinChatRoom(chatRoomName string, user string) error {
	chatRoom, err := service.chatRoomRepo.GetChatRoom(chatRoomName)

	if err != nil {
		log.Println("Error getting chat room: ", err)
		return err
	}

	chatRoom.AddUser(user)
	_, err = service.chatRoomRepo.UpdateChatRoom(chatRoom)

	return err
}

func (service *chatRoomServiceImpl) LeaveChatRoom(chatRoomName string, user string) error {
	chatRoom, err := service.chatRoomRepo.GetChatRoom(chatRoomName)

	if err != nil {
		log.Println("Error getting chat room: ", err)
		return err
	}

	chatRoom.RemoveUser(user)

	if len(chatRoom.GetUsers()) == 0 {
		err = service.chatRoomRepo.DeleteChatRoom(chatRoomName)
		return err
	} else {
		_, err = service.chatRoomRepo.UpdateChatRoom(chatRoom)
	}

	return err
}
