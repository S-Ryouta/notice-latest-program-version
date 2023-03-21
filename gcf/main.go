package gcf

import (
	"context"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/infrastructure/db"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/usecase/version"
	"github.com/go-redis/redis/v8"
	"os"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func CheckAndUpdateVersionHandler(ctx context.Context, m PubSubMessage) error {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   1, // use default DB
	})

	redisVersionRepo := db.NewRedisVersionRepository(redisClient)
	versionInteractor := version.NewVersionInteractor(redisVersionRepo)
	versionInteractor.CheckAndUpdateVersion()
	return nil
}
