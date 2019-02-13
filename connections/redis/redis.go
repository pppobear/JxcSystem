package redis

import (
	"github.com/go-redis/redis"
	"pppobear.cn/jxc-backend/config"
	"time"
)

type ClientType struct {
	RedisCon *redis.Client
}

var Client ClientType

func init() {
	Client.RedisCon = redis.NewClient(&redis.Options{
		Addr:     config.GetEnv().RedisIP + ":" + config.GetEnv().RedisPort,
		Password: config.GetEnv().RedisPassword,
		DB:       config.GetEnv().RedisDb,
	})
}

func (client *ClientType) Set(key string, value interface{}, expiration time.Duration) *redis.Client {
	err := client.RedisCon.Set(key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	return (*client).RedisCon
}

func (client *ClientType) Get(key string) (string, *redis.Client) {
	val, err := (*client).RedisCon.Get(key).Result()

	if err == redis.Nil {
		return "", (*client).RedisCon
	}

	if err != nil {
		panic(err)
	}

	return val, (*client).RedisCon
}

func (client *ClientType) Del(key string) *redis.Client {
	_, err := client.RedisCon.Del(key).Result()
	if err != nil {
		panic(err)
	}
	return (*client).RedisCon
}
