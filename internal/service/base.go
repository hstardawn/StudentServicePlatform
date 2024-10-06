package service

import (
	"StudentServicePlatform/internal/dao"
	"StudentServicePlatform/internal/global"
	"StudentServicePlatform/internal/pkg/redis"
	"context"
	"time"

	"gorm.io/gorm"
)

var (
	ctx = context.Background()
	d  *dao.Dao
)

func ServiceInit(db *gorm.DB) {
	d =dao.New(db)
}

func GetConfigUrl() string {
	url := GetRedis("url")
	if url == "" {
		url = global.Config.GetString("url.host")
		SetRedis("url", url)
	}
	return url
}

func GetConfigKey() string {
	key := GetRedis("key")
	if key == "" {
		key = global.Config.GetString("key")
		SetRedis("key", key)
	}
	return key
}

func SetRedis(key string, value string) bool {
	t := int64(900)
	expire := time.Duration(t) * time.Second
	if err := redis.RedisClient.Set(ctx, key, value, expire).Err(); err != nil {
		return false
	}
	return true
}

func GetRedis(key string) string {
	result, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}