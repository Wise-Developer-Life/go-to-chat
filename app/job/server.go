package job

import (
	"github.com/hibiken/asynq"
	"go-to-chat/app/utility"
	"log"
)

func CreateJobServer() *asynq.Server {
	redisConfig, err := utility.GetRedisConfig(utility.TypeJobCache)

	if err != nil {
		log.Println("Error getting redis config: ", err)
		return nil
	}

	mux := asynq.NewServeMux()
	mux.Handle(string(TypeFindMatch), NewMatchingProcessor())

	server := asynq.NewServer(asynq.RedisClientOpt{
		Addr:     redisConfig.Address(),
		DB:       redisConfig.DB,
		Password: redisConfig.Password,
	}, asynq.Config{})

	go func() {
		if err := server.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	}()

	return server
}
