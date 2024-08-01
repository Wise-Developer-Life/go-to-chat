package service

import (
	"go-to-chat/app/job"
	"go-to-chat/app/utility"
	"log"
	"time"
)

type MatchService interface {
	CreateNewMatchTask(user string) error
}

type matchServiceImpl struct {
	jobClient *job.Client
}

func NewMatchService() MatchService {
	return &matchServiceImpl{
		jobClient: job.GetClientInstance(),
	}
}

func (service *matchServiceImpl) CreateNewMatchTask(user string) error {
	redisClient, err := utility.GetRedisClient(utility.TypeDefault)

	if err != nil {
		log.Println("Error getting redis client's instance: ", err)
		return err
	}

	currentTime := time.Now().Unix()
	err = redisClient.ZAdd("matching_pool", user, float64(currentTime))

	// add matching task for user
	err = service.jobClient.Enqueue(
		job.TypeFindMatch,
		&job.FindMatchJobPayload{
			User: user,
		},
		&job.EnqueueConfig{
			Timout: time.Minute * 5,
		})

	if err != nil {
		log.Println("Error enqueuing task: ", err)
		return err
	}
	return nil
}
