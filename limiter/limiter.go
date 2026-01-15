package limiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	client *redis.Client
}

func NewRateLimiter(client *redis.Client) *RateLimiter {
	return &RateLimiter{
		client: client,
	}
}

func (r *RateLimiter) Allow(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	luaScript := `
		local current = redis.call("INCR", KEYS[1])
		if tonumber(current) == 1 then
			redis.call("EXPIRE", KEYS[1], ARGV[1])
		end
		return current
	`

	result, err := r.client.Eval(ctx, luaScript, []string{key}, int(window.Seconds())).Int()
	if err != nil {
		return false, err
	}

	return result <= limit, nil
}
