package redis

import (
	"github.com/go-redis/redis"

	"github.com/ihahoo/go-api-lib/config"
	"github.com/ihahoo/go-api-lib/log"
)

// Conn 连接redis
func Conn(opt *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(opt)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Client 返回redis客户端
func Client(db int) *redis.Client {
	host := config.GetString("redis.host")
	port := config.GetString("redis.port")
	password := config.GetString("redis.password")

	client, err := Conn(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
	if err != nil {
		log.GetLog().Fatal(err)
	}
	return client
}
