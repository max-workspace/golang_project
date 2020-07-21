package goredisadapter

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type adapter struct {
	redisClient *redis.Client
}

var (
	adapterInstance *adapter
	loadAdapterOnce sync.Once
)

// New new redis client adapter
func New(addr, password string, db int) *adapter {
	loadAdapterOnce.Do(func() {
		redisClientInstance := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
		adapterInstance = &adapter{}
		adapterInstance.redisClient = redisClientInstance
	})
	return adapterInstance

}

func (redis *adapter) Close() error {
	return redis.redisClient.Close()
}

func (redis *adapter) Exists(keys ...string) error {
	return redis.redisClient.Exists(keys...).Err()
}

func (redis *adapter) Expire(key string, expiration time.Duration) error {
	return redis.redisClient.Expire(key, expiration).Err()
}

func (redis *adapter) Set(key string, value interface{}, expiration time.Duration) error {
	return redis.redisClient.Set(key, value, expiration).Err()
}

func (redis *adapter) Get(key string) (string, error) {
	return redis.redisClient.Get(key).Result()
}

func (redis *adapter) Del(keys ...string) error {
	return redis.redisClient.Del(keys...).Err()
}
