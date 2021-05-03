package redis

import (
	"context"
	"time"
)

// GetBool .
func (rds *Redis) GetBool(ctx context.Context, key string) (bool, error) {
	b, err := rds.client.Get(ctx, key).Bool()
	if err != nil {
		return false, err
	}

	return b, nil
}

// SetBool .
func (rds *Redis) SetBool(ctx context.Context, key string, value bool) error {
	return rds.client.Set(ctx, key, value, 0).Err()
}

// SetExpBool .
func (rds *Redis) SetExpBool(ctx context.Context, key string, value bool, expiration time.Duration) error {
	return rds.client.Set(ctx, key, value, expiration).Err()
}
