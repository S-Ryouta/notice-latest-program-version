package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/domain/entity"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisVersionRepository struct {
	client redis.Cmdable
}

func NewRedisVersionRepository(client redis.Cmdable) *RedisVersionRepository { // 修正
	return &RedisVersionRepository{
		client: client,
	}
}

func (r *RedisVersionRepository) GetVersion() (*entity.Version, error) {
	ctx := context.Background()
	versionJSON, err := r.client.Get(ctx, "golang_version").Result()
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

	err = r.client.Set(ctx, "golang_version", versionJSON, time.Hour*24*31).Err()
	if err != nil {
		return err
	}

	return nil
}
