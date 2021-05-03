package redis

import (
	"context"
	"time"
)

// GetStr .
func (rds *Redis) GetStr(ctx context.Context, key string) (string, error) {
	strCmd := rds.client.Get(ctx, key)

	return strCmd.String(), strCmd.Err()
}

// SetStr .
func (rds *Redis) SetStr(ctx context.Context, key string, value string) error {
	return rds.client.Set(ctx, key, value, 0).Err()
}

// SetExpStr .
func (rds *Redis) SetExpStr(ctx context.Context, key string, value string, expiration time.Duration) error {
	return rds.client.Set(ctx, key, value, expiration).Err()
}
