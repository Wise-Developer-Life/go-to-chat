package job

import (
	"context"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"go-to-chat/app/chat"
	"go-to-chat/app/user"
	"go-to-chat/app/utility"
	"log"
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
	hub := chat.GetHubInstance()

	currentUser := payload.Content.User
	isUserInMatchingPool, _ := redisClient.ZExists("matching_pool", currentUser)

	if !isUserInMatchingPool || !hub.IsUserConnected(currentUser) {
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
		if user != currentUser && hub.IsUserConnected(user) {
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

	//TODO: implement chat room creation
	room := chat.CreateRoom()
	clientCurrentUser := hub.GetClientFromUsername(currentUser)
	clientMatchedUser := hub.GetClientFromUsername(matchedUser)
	room.Join(clientCurrentUser, clientMatchedUser)
	log.Println(fmt.Sprintf("Client %s joined room", currentUser))
	log.Println(fmt.Sprintf("Client %s joined room", matchedUser))

	matchedResCurrentUser := &chat.SocketMatchUserResponse{
		MatchedUser: &user.UserResponse{
			ID:    matchedUserEntity.ID,
			Name:  matchedUserEntity.Name,
			Email: matchedUserEntity.Email,
		},
	}
	clientCurrentUser.Send(&chat.SocketResponse[*chat.SocketMatchUserResponse]{
		Event: chat.SocketEventMatched,
		Data:  matchedResCurrentUser,
	})

	matchedResMatchedUser := &chat.SocketMatchUserResponse{
		MatchedUser: &user.UserResponse{
			ID:    currentUserEntity.ID,
			Name:  currentUserEntity.Name,
			Email: currentUserEntity.Email,
		},
	}
	clientMatchedUser.Send(&chat.SocketResponse[*chat.SocketMatchUserResponse]{
		Event: chat.SocketEventMatched,
		Data:  matchedResMatchedUser,
	})

	_ = redisClient.ZRemove("matching_pool", currentUser)
	_ = redisClient.ZRemove("matching_pool", matchedUser)
	log.Println("process finished")
	return nil
}
