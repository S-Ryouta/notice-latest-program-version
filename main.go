package notice_latest_program_version

import (
	"context"
	"github.com/S-Ryouta/notice-latest-program-version/infrastructure/db"
	"github.com/S-Ryouta/notice-latest-program-version/infrastructure/version"
	usecaseVersion "github.com/S-Ryouta/notice-latest-program-version/usecase/version"
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

	versionGetter := version.NewDefaultVersionGetter()
	redisVersionRepo := db.NewRedisVersionRepository(redisClient)
	versionInteractor := usecaseVersion.NewVersionInteractor(redisVersionRepo, versionGetter)
	versionInteractor.CheckAndUpdateVersion()
	return nil
}
