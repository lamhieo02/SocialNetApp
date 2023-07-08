package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	redis *redis.Client
}

func NewCache(redis *redis.Client) *cache {
	return &cache{redis: redis}
}

func (c *cache) Set(ctx context.Context, key string, value interface{}) error {
	err := c.redis.Set(ctx, key, value, 1*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *cache) Get(ctx context.Context, key string) (*string, error) {
	val, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}
