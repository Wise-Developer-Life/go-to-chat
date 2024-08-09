package job

import (
	"context"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"go-to-chat/app/chat"
	"go-to-chat/app/match/response"
	"go-to-chat/app/socket"
	"go-to-chat/app/user"
	"go-to-chat/app/utility"
	"log"
	"sort"
	"time"
)

type MatchingProcessor struct{}

func NewMatchingProcessor() *MatchingProcessor {
	return &MatchingProcessor{}
}

// ProcessTask TODO: implement this method
var userService = user.NewUserService(user.NewUserRepository())

func (p *MatchingProcessor) ProcessTask(_ context.Context, task *asynq.Task) error {
	log.Println("process started")
	// decode payload to json
	payload, err := DecodeJobPayload[*FindMatchJobPayload](task.Payload())

	if err != nil {
		return err
	}

	redisClient, _ := utility.GetRedisClient(utility.TypeDefault)
	hub := socket.GetHubInstance()

	currentUser := payload.Content.User
	isUserInMatchingPool, _ := redisClient.ZExists("matching_pool", currentUser)

	if !isUserInMatchingPool || !hub.IsClientConnected(currentUser) {
		log.Println(fmt.Sprintf("User %s is not in matching pool or not connected", currentUser))
		return nil
	}

	// FIXME: simulate 3 seconds wait
	time.Sleep(3 * time.Second)

	// massive task
	userList, _ := redisClient.ZRange("matching_pool", 0, 10, false)
	matchedUser := ""
	for _, user := range userList {
		user := user.(string)
		if user != currentUser && hub.IsClientConnected(user) {
			matchedUser = user
			break
		}
	}

	//FIXME: add retry mechanism
	if matchedUser == "" {
		log.Println("No matched user found for: " + currentUser)
		return errors.New("no matched user found")
	}

	log.Println("Matched user: " + matchedUser)

	currentUserEntity, err := userService.GetUserByEmail(currentUser)
	if err != nil {
		return err
	}

	matchedUserEntity, err := userService.GetUserByEmail(matchedUser)
	if err != nil {
		return err
	}

	// refactor create chat room

	chatRoomService := chat.GetChatRoomServiceInstance()

	if chatRoomService == nil {
		return errors.New("error getting chat room service instance")
	}

	room, err := chatRoomService.CreateChatRoom(generateRoomID(currentUser, matchedUser))

	if err != nil {
		return err
	}

	err = chatRoomService.JoinChatRoom(room.ID, currentUser)
	err = chatRoomService.JoinChatRoom(room.ID, matchedUser)

	log.Println(fmt.Sprintf("Client %s and %s joined room %s", currentUser, matchedUser, room.ID))

	clientCurrentUser := hub.GetClient(currentUser)
	clientMatchedUser := hub.GetClient(matchedUser)
	matchedResCurrentUser := &response.MatchUserResponse{
		MatchedUser: user.NewUserResponse(matchedUserEntity),
	}
	clientCurrentUser.Send(
		socket.NewSocketMessage(socket.SocketEventMatched, matchedResCurrentUser),
	)

	matchedResMatchedUser := &response.MatchUserResponse{
		MatchedUser: user.NewUserResponse(currentUserEntity),
	}
	clientMatchedUser.Send(
		socket.NewSocketMessage(socket.SocketEventMatched, matchedResMatchedUser),
	)

	_ = redisClient.ZRemove("matching_pool", currentUser)
	_ = redisClient.ZRemove("matching_pool", matchedUser)
	log.Println("process finished")
	return nil
}

func generateRoomID(user1 string, user2 string) string {
	users := []string{user1, user2}
	sort.Strings(users)
	return users[0] + ":" + users[1]
}
