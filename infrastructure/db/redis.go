package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/S-Ryouta/notice-latest-program-version/domain/entity"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisVersionRepository struct {
	client redis.Cmdable
}

func NewRedisVersionRepository(client redis.Cmdable) *RedisVersionRepository {
	return &RedisVersionRepository{
		client: client,
	}
}

func (r *RedisVersionRepository) GetVersion(language string) (*entity.Version, error) {
	ctx := context.Background()
	versionJSON, err := r.client.Get(ctx, fmt.Sprintf("%s_version", language)).Result()
	var version entity.Version

	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
		return &version, nil
	case err != nil:
		return nil, err
	}

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

	err = r.client.Set(ctx, fmt.Sprintf("%s_version", version.ID), versionJSON, time.Hour*24*31).Err()
	if err != nil {
		return err
	}

	return nil
}
