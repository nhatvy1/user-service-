package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"user-service/pkg/globals"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string, dest any) error
	Set(ctx context.Context, key string, value any, time time.Duration) error
	Clear(ctx context.Context, key string) error
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisCache() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", globals.Config.Redis.Host, globals.Config.Redis.Port),
		Password: globals.Config.Redis.Password,
		DB:       0,
	})

	return &RedisClient{
		client: rdb,
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, b, time.Duration(ttl)*time.Second).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string, dest any) error {
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

func (r *RedisClient) Clear(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
