package redis

import (
	"context"
	"time"
)

// GetByte .
func (rds *Redis) GetByte(ctx context.Context, key string) ([]byte, error) {
	b, err := rds.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	return b, nil
}

// SetByte .
func (rds *Redis) SetByte(ctx context.Context, key string, value []byte) error {
	return rds.client.Set(ctx, key, value, 0).Err()
}

// SetExpByte .
func (rds *Redis) SetExpByte(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	return rds.client.Set(ctx, key, value, expiration).Err()
}
