package job

import (
	"github.com/hibiken/asynq"
	"go-to-chat/app/utility"
	"log"
	"sync"
	"time"
)

type ClientInterface interface {
	Enqueue(jobType Type, content any) error
	Close() error
}

type EnqueueConfig struct {
	Timout time.Duration
}

var (
	instance *Client
	once     sync.Once
)

type Client struct {
	client *asynq.Client
}

func newJobClient(client *asynq.Client) *Client {
	return &Client{client: client}
}

func GetClientInstance() *Client {
	redisConfig, err := utility.GetRedisConfig(utility.TypeJobCache)

	if err != nil {
		log.Println("Error getting redis config: ", err)
		return nil
	}

	once.Do(func() {
		queueClient := asynq.NewClient(asynq.RedisClientOpt{
			Addr:     redisConfig.Address(),
			DB:       redisConfig.DB,
			Password: redisConfig.Password,
		})
		instance = newJobClient(queueClient)
	})
	return instance
}

func (c *Client) Enqueue(jobType Type, content any, config *EnqueueConfig) error {
	payload := &Common[any]{
		CreatedAt: time.Now().Unix(),
		Content:   content,
	}

	payloadBytes, err := EncodeJobPayload(payload)

	if err != nil {
		log.Fatalf("could not marshal payload: %v", err)
		return err
	}

	_, err = c.client.Enqueue(
		asynq.NewTask(string(jobType), payloadBytes),
		asynq.Timeout(config.Timout),
	)
	return err
}

func (c *Client) Close() error {
	return c.client.Close()
}
