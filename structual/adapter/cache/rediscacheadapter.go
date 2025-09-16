package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisCacheAdapter struct {
	client *redis.Client
}

func NewRedisCacheAdapter(client *redis.Client) *RedisCacheAdapter {
	return &RedisCacheAdapter{client: client}
}

func (c *RedisCacheAdapter) Put(key string, value string) error {
	c.client.Set(context.Background(), key, value, 0)
	return nil
}

func (c *RedisCacheAdapter) Get(key string) (string, error) {
	value, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
func (c *RedisCacheAdapter) Remove(key string) {
	c.client.Del(context.Background(), key)
}
