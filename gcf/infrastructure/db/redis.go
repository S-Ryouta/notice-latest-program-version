package db

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"notice-latest-program-version/domain/entity"
	"time"
)

type RedisVersionRepository struct {
	client *redis.Client
}

func NewRedisVersionRepository(client *redis.Client) *RedisVersionRepository {
	return &RedisVersionRepository{
		client: client,
	}
}

func (r *RedisVersionRepository) GetVersion() (*entity.Version, error) {
	ctx := context.Background()
	versionJSON, err := r.client.Get(ctx, "golang_version").Result()
	if err != nil {
		return nil, err
	}

	var version entity.Version
	err = json.Unmarshal([]byte(versionJSON), &version)
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func (r *RedisVersionRepository) SaveVersion(version *entity.Version) error {
	ctx := context.Background()
	versionJSON, err := json.Marshal(version)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, "golang_version", versionJSON, time.Hour*24*7).Err()
	if err != nil {
		return err
	}

	return nil
}
