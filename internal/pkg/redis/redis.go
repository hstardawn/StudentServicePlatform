package redis

import (
	"StudentServicePlatform/internal/global"

	"github.com/go-redis/redis/v8"
)

type RedisInfoConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

func GetRedisClient(info RedisInfoConfig) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     info.Host + ":" + info.Port,
		Password: info.Password,
		DB:       info.DB,
	})
	return redisClient
}

func DefaultRedisConfig() RedisInfoConfig {
	return RedisInfoConfig{
		Host:     "localhost",
		Port:     "6379",
		DB:       0,
		Password: "",
	}
}
func getConfig() RedisInfoConfig {
	Info := RedisInfoConfig{
		Host:     "localhost",
		Port:     "6379",
		DB:       0,
		Password: "",
	}
	if global.Config.IsSet("redis.host") {
		Info.Host = global.Config.GetString("redis.host")
	}
	if global.Config.IsSet("redis.port") {
		Info.Port = global.Config.GetString("redis.port")
	}
	if global.Config.IsSet("redis.db") {
		Info.DB = global.Config.GetInt("redis.db")
	}
	if global.Config.IsSet("redis.pass") {
		Info.Password = global.Config.GetString("redis.pass")
	}
	return Info
}
