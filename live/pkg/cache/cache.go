package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	keyPrefix = "doogeun.live."
)

type Cache interface {
	GetJSON(ctx context.Context, key string) (json.RawMessage, error)
	SetJSON(ctx context.Context, key string, value json.RawMessage, ttl ...time.Duration) error
}

type cache struct {
	client *redis.Client
}

func New(host string, port int) Cache {
	cache := &cache{}
	cache.client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})
	return cache
}

func (c *cache) GetJSON(ctx context.Context, key string) (json.RawMessage, error) {
	value, err := c.client.Get(ctx, keyPrefix+key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	var msg json.RawMessage
	err = json.Unmarshal([]byte(value), &msg)
	return msg, err
}

func (c *cache) SetJSON(ctx context.Context, key string, value json.RawMessage, ttl ...time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var ttlV time.Duration = 0
	if len(ttl) > 0 {
		ttlV = ttl[0]
	}
	return c.client.Set(ctx, keyPrefix+key, data, ttlV).Err()
}
