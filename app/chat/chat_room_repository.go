package chat

import (
	"go-to-chat/app/model"
	"go-to-chat/app/utility"
	"log"
)

type RoomRepository interface {
	CreateChatRoom(chatRoomName string) (*model.ChatRoom, error)
	GetChatRoom(chatRoomName string) (*model.ChatRoom, error)
	UpdateChatRoom(chatRoom *model.ChatRoom) (*model.ChatRoom, error)
	DeleteChatRoom(chatRoomName string) error
}

type RoomRepositoryRedisImpl struct {
	redisClient utility.RedisClient
}

func NewChatRoomRepository() RoomRepository {
	redisClient, err := utility.GetRedisClient(utility.TypeDefault)

	if err != nil {
		log.Println("Error getting redis client's instance: ", err)
		return nil
	}

	return &RoomRepositoryRedisImpl{
		redisClient: redisClient,
	}
}

func (repo *RoomRepositoryRedisImpl) GetChatRoom(chatRoomName string) (*model.ChatRoom, error) {
	chatRoomStr, err := repo.redisClient.Get(chatRoomName)

	if err != nil {
		return nil, err
	}

	var chatRoom model.ChatRoom
	err = chatRoom.UnmarshalBinary([]byte(chatRoomStr))

	if err != nil {
		return nil, err
	}

	return &chatRoom, nil
}

func (repo *RoomRepositoryRedisImpl) CreateChatRoom(chatRoomName string) (*model.ChatRoom, error) {
	chatRoom := model.NewChatRoom(chatRoomName)

	err := repo.redisClient.Set(chatRoomName, chatRoom, 0)

	if err != nil {
		return nil, err
	}

	return chatRoom, nil
}

func (repo *RoomRepositoryRedisImpl) UpdateChatRoom(chatRoom *model.ChatRoom) (*model.ChatRoom, error) {
	err := repo.redisClient.Set(chatRoom.ID, chatRoom, 0)

	if err != nil {
		return nil, err
	}

	return chatRoom, nil
}

func (repo *RoomRepositoryRedisImpl) DeleteChatRoom(chatRoomName string) error {
	return repo.redisClient.Del(chatRoomName)
}
