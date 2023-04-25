package storage

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)
var Redis *redis.Client

func init() {
	Redis := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		log.Fatalf("Error while pinging redis server %v",err)
	}
	fmt.Println(pong)
	redisSize := Redis.DBSize().Val()
	fmt.Println(redisSize)
}