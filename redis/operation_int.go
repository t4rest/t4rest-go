package redis

import (
	"context"
	"time"
)

// GetInt .
func (rds *Redis) GetInt(ctx context.Context, key string) (int, error) {
	b, err := rds.client.Get(ctx, key).Int()
	if err != nil {
		return 0, err
	}

	return b, nil
}

// SetInt .
func (rds *Redis) SetInt(ctx context.Context, key string, value int) error {
	return rds.client.Set(ctx, key, value, 0).Err()
}

// SetExpInt .
func (rds *Redis) SetExpInt(ctx context.Context, key string, value int, expiration time.Duration) error {
	return rds.client.Set(ctx, key, value, expiration).Err()
}
