package main

import (
	"github.com/go-redis/redis/v7"
)

func redisConnect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		PoolSize:     10,
		MinIdleConns: 5,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return client, nil
}
func main() {

}
