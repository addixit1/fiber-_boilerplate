package redis

import (
	"context"
	"log"
	"time"

	"github.com/addixit1/fiber-boilerplate/internal/config"
	"github.com/addixit1/fiber-boilerplate/internal/utils"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:         config.Config.RedisURI,
		Password:     "", // set if required
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     20,
		MinIdleConns: 5,
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("redis connection failed: %v", err)
	}
	utils.LogDatabase("Redis connected successfully on localhost:6379")
}

func Publish(channel string, message string) error {
	return Client.Publish(Ctx, channel, message).Err()
}

func Subscribe(channel string) {
	sub := Client.Subscribe(Ctx, channel)
	ch := sub.Channel()

	log.Println("Subscribed to:", channel)

	for msg := range ch {
		log.Println("Received:", msg.Payload)
	}
}

func Set(key string, value string, ttl time.Duration) error {
	return Client.Set(Ctx, key, value, ttl).Err()
}

func Get(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}
