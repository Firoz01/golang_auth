package cmd

import (
	"context"
	"fmt"
	"github.com/go-frame/config"
	"github.com/go-frame/internals/lib/logger"
	"github.com/redis/go-redis/v9"
	"time"
)

func setupRedis(conf *config.Config, appLogger logger.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check if Redis is connected
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		appLogger.Error("Failed to connect to Redis", "error", err)
		return nil, fmt.Errorf("could not connect to Redis: %v", err)
	}

	appLogger.Info("Successfully connected to Redis")
	return client, nil
}
