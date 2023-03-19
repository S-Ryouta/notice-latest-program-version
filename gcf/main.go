package main

import (
	"context"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/infrastructure/db"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/usecase/version"
	"github.com/go-redis/redis/v8"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func CheckAndUpdateVersionHandler(ctx context.Context, m PubSubMessage) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisVersionRepo := db.NewRedisVersionRepository(redisClient)
	versionInteractor := version.NewVersionInteractor(redisVersionRepo)
	versionInteractor.CheckAndUpdateVersion()
}
