package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	// Client client
	Client *redis.Client
)

// Init redis client
func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("app.redis.localDemo.addr"),
		Password: viper.GetString("app.redis.localDemo.password"),
		DB:       viper.GetInt("app.redis.localDemo.DB"),
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
