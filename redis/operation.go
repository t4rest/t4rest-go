package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

// Exists .
func (rds *Redis) Exists(ctx context.Context, key string) (bool, error) {
	res, err := rds.client.Exists(ctx, key).Result()
	exists := res > 0

	return exists, err
}

// Delete .
func (rds *Redis) Delete(ctx context.Context, key string) error {
	return rds.client.Del(ctx, key).Err()
}

// Incr .
func (rds *Redis) Incr(ctx context.Context, key string) (int64, error) {
	v := rds.client.Incr(ctx, key)

	return v.Val(), v.Err()
}

// Expire .
func (rds *Redis) Expire(ctx context.Context, key string, exp time.Duration) error {
	return rds.client.Expire(ctx, key, exp).Err()
}

// CheckTTL .
func (rds *Redis) CheckTTL(ctx context.Context, key string) error {
	res := rds.client.TTL(ctx, key)
	if res.Err() != nil {
		return res.Err()
	}

	const expIsNotSet = -1
	if res.Val().Nanoseconds() == expIsNotSet {
		return errors.New("ttl is not set")
	}

	return nil
}
