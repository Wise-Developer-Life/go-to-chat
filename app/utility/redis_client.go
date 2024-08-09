package utility

//TODO: write unit tests
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type ConnectionType string

const (
	TypeDefault  ConnectionType = "default"
	TypeJobCache ConnectionType = "job_cache"
)

type ConnectionConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func (c *ConnectionConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

var configSet = map[ConnectionType]*ConnectionConfig{
	"default": {
		Host:     "localhost",
		Port:     6378,
		Password: "",
		DB:       0,
	},
	"job_cache": {
		Host:     "localhost",
		Port:     6378,
		Password: "",
		DB:       1,
	},
}

var ctx = context.Background()

var mapRedisClient = make(map[ConnectionType]RedisClient)

type RedisClient interface {
	Set(key string, value any, ttl time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	ZAdd(key string, value any, score float64) error
	ZRange(key string, start, end int64, reverse bool) ([]any, error)
	ZRemove(key string, value any) error
	ZExists(key string, value any) (bool, error)
}

type redisClientImpl struct {
	connection ConnectionType
	client     *redis.Client
}

func newRedisClientImpl(connection ConnectionType, config *ConnectionConfig) (*redisClientImpl, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address(),
		Password: config.Password,
		DB:       config.DB,
	})

	if rdb == nil {
		return nil, errors.New("failed to create redis client")
	}

	return &redisClientImpl{
		connection: connection,
		client:     rdb,
	}, nil
}

func GetRedisConfig(connection ConnectionType) (*ConnectionConfig, error) {
	if redisConfig, ok := configSet[connection]; ok {
		return redisConfig, nil
	}
	return nil, errors.New("connection not found")
}

func GetRedisClient(connection ConnectionType) (RedisClient, error) {
	if _, ok := configSet[connection]; !ok {
		return nil, errors.New("connection not found")
	}

	if _, ok := mapRedisClient[connection]; !ok {
		newCreatedClient, err := newRedisClientImpl(connection, configSet[connection])

		if err != nil {
			return nil, err
		}

		mapRedisClient[connection] = newCreatedClient
	}

	return mapRedisClient[connection], nil
}

func (r *redisClientImpl) Set(key string, value any, ttl time.Duration) error {
	status := r.client.Set(ctx, key, value, ttl)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (r *redisClientImpl) Get(key string) (string, error) {
	status := r.client.Get(ctx, key)
	if status.Err() != nil {
		return "", status.Err()
	}
	return status.Val(), nil
}

func (r *redisClientImpl) Del(key string) error {
	status := r.client.Del(ctx, key)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func (r *redisClientImpl) ZAdd(key string, value any, score float64) error {
	jsonByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	member := &redis.Z{
		Score:  score,
		Member: jsonByte,
	}

	status := r.client.ZAdd(ctx, key, *member)

	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (r *redisClientImpl) ZRange(key string, start, end int64, reverse bool) ([]any, error) {
	var status *redis.StringSliceCmd
	if reverse {
		status = r.client.ZRevRange(ctx, key, start, end)
	} else {
		status = r.client.ZRange(ctx, key, start, end)
	}

	if status.Err() != nil {
		return nil, status.Err()
	}

	ret := status.Val()

	var result []any
	for _, val := range ret {
		var payload any
		err := json.Unmarshal([]byte(val), &payload)
		if err != nil {
			return nil, err
		}

		result = append(result, payload)
	}

	return result, nil
}

func (r *redisClientImpl) ZRemove(key string, value any) error {
	jsonByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	status := r.client.ZRem(ctx, key, jsonByte)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (r *redisClientImpl) ZExists(key string, value any) (bool, error) {
	jsonByte, err := json.Marshal(value)
	if err != nil {
		return false, err
	}

	status := r.client.ZScore(ctx, key, string(jsonByte))
	if status.Err() != nil {
		return false, status.Err()
	}
	return true, nil
}
