package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/yerassyldanay/invest/utils/config"
)

var redisClient *redis.Client

func EstablishConnectionWithRedis(opts config.Config) (*redis.Client, error) {
	// redis
	redisClientTemp := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", opts.RedisHost, opts.RedisPort),
		DB:   0,
	})

	// check connection
	cmd := redisClientTemp.Ping()
	if cmd.Err() != nil {
		return nil, fmt.Errorf("failed to ping redis. err: %v", cmd.Err())
	}

	redisClient = redisClientTemp
	return redisClient, nil
}

func GetRedis() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	// get variables
	opts, err := config.LoadConfig("../environment/local.env")
	if err != nil {
		fmt.Printf("failed to get env variables. err: %v\n", err)
	}

	// establish otherwise
	c, err := EstablishConnectionWithRedis(opts)
	if err != nil {
		fmt.Printf("failed to establish connection with redis. err: %v\n", err)
	}

	redisClient = c
	return c
}

