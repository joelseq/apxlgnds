package cache

import (
	"context"
	"errors"
	"time"

	"github.com/joelseq/apxlgnds/internal/types"
	"github.com/redis/go-redis/v9"
)

type cache struct {
	client *redis.Client
}

var ErrCacheEmpty = errors.New("cache is empty")

type Cacher interface {
	GetResult(ctx context.Context) (*types.CalendarEventsResponse, error)
	SetResult(ctx context.Context, response *types.CalendarEventsResponse) error
}

const (
	cacheKey = "calendar_events"
)

func NewCache(addr, password string) Cacher {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &cache{
		client: client,
	}
}

func (c *cache) GetResult(ctx context.Context) (*types.CalendarEventsResponse, error) {
	val, err := c.client.Get(ctx, cacheKey).Bytes()
	if err == redis.Nil {
		return nil, ErrCacheEmpty
	}
	if err != nil && err != redis.Nil {
		return nil, err
	}

	return types.DecodeResponse(val)
}

func (c *cache) SetResult(ctx context.Context, response *types.CalendarEventsResponse) error {
	val, err := types.EncodeResponse(response)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, cacheKey, val, 1*time.Hour).Err()
}
