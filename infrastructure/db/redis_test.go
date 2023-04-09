package db_test

import (
	"context"
	"errors"
	"github.com/S-Ryouta/notice-latest-program-version/domain/entity"
	"github.com/S-Ryouta/notice-latest-program-version/infrastructure/db"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRedisClient struct {
	redis.Cmdable
	mock.Mock
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	return args.Get(0).(*redis.StringCmd)
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	return args.Get(0).(*redis.StatusCmd)
}

func TestRedisVersionRepository_GetVersion(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := new(MockRedisClient)
		repo := db.NewRedisVersionRepository(client)
		language := "golang"

		versionJSON := `{"version": "1.0.0"}`
		client.On("Get", mock.Anything, "golang_version").Return(redis.NewStringResult(versionJSON, nil))

		expected := &entity.Version{
			Version: "1.0.0",
		}

		version, err := repo.GetVersion(language)
		assert.NoError(t, err)
		assert.Equal(t, expected, version)
		client.AssertExpectations(t)
	})

	t.Run("key does not exist", func(t *testing.T) {
		client := new(MockRedisClient)
		repo := db.NewRedisVersionRepository(client)
		language := "golang"

		client.On("Get", mock.Anything, "golang_version").Return(redis.NewStringResult("", redis.Nil))

		version, err := repo.GetVersion(language)
		assert.NoError(t, err)
		assert.NotNil(t, version)
		assert.Empty(t, version.Version)
		client.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		client := new(MockRedisClient)
		repo := db.NewRedisVersionRepository(client)
		language := "golang"

		client.On("Get", mock.Anything, "golang_version").Return(redis.NewStringResult("", errors.New("some error")))

		_, err := repo.GetVersion(language)
		assert.Error(t, err)
		client.AssertExpectations(t)
	})
}

func TestRedisVersionRepository_SaveVersion(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := new(MockRedisClient)
		repo := db.NewRedisVersionRepository(client)

		version := &entity.Version{
			ID:      "golang",
			Version: "1.0.0",
		}

		client.On("Set", mock.Anything, "golang_version", mock.Anything, time.Hour*24*31).Return(redis.NewStatusResult("", nil))

		err := repo.SaveVersion(version)
		assert.NoError(t, err)
		client.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		client := new(MockRedisClient)
		repo := db.NewRedisVersionRepository(client)

		version := &entity.Version{
			ID:      "golang",
			Version: "1.0.0",
		}

		client.On("Set", mock.Anything, "golang_version", mock.Anything, time.Hour*24*31).Return(redis.NewStatusResult("", errors.New("some error")))

		err := repo.SaveVersion(version)
		assert.Error(t, err)
		client.AssertExpectations(t)
	})
}
