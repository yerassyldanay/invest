package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

var redisClient *redis.Client

func New_redis_connection(db_num int) *redis.Client {
	fmt.Println("Connecting to Redis...")

	/*
		here redis uri will be created
	 */
	var redisPassword = os.Getenv("REDIS_PASSWORD")
	var redisHost = os.Getenv("POSTGRES_HOST")
	var redisPort = os.Getenv("REDIS_PORT")

	rClient := redis.NewClient(&redis.Options{
		Addr:			redisHost + ":" + redisPort,
		Password:	 	redisPassword,
		DB: 			db_num,
	})

	_, err := rClient.Ping().Result()

	if err != nil {
		fmt.Println("Redis connection failed ...")
	}

	redisClient = rClient
	return rClient
}

func GetRedis() (*redis.Client) {
	return redisClient
}
