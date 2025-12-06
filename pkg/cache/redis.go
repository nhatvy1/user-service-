package cache

import "time"

type RedisCacheService interface {
	Get(key string, value any) (string, error)
	Set(key string, value any, time time.Duration) error
	Clear(key string) error
}
