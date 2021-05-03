package redis

import (
	"context"
	"time"

	rds "github.com/go-redis/redis/v8"
)

// Cacher base redis interface
type Cacher interface {
	GetByte(ctx context.Context, key string) ([]byte, error)
	SetByte(ctx context.Context, key string, value []byte) error
	SetExpByte(ctx context.Context, key string, value []byte, expiration time.Duration) error

	GetInt(ctx context.Context, key string) (int, error)
	SetInt(ctx context.Context, key string, value int) error
	SetExpInt(ctx context.Context, key string, value int, expiration time.Duration) error

	GetStr(ctx context.Context, key string) (string, error)
	SetStr(ctx context.Context, key string, value string) error
	SetExpStr(ctx context.Context, key string, value string, expiration time.Duration) error

	GetBool(ctx context.Context, key string) (bool, error)
	SetBool(ctx context.Context, key string, value bool) error
	SetExpBool(ctx context.Context, key string, value bool, expiration time.Duration) error

	Exists(ctx context.Context, key string) (bool, error)
	Delete(ctx context.Context, key string) error

	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error

	Close() error
}

// Redis .
type Redis struct {
	client *rds.Client
}

// New create new pool redis connections
func New(cfg Conf) (*Redis, error) {
	r, err := NewClient(cfg)

	return &Redis{client: r}, err
}

// NewClient .
func NewClient(cfg Conf) (*rds.Client, error) {
	var client *rds.Client

	if len(cfg.SentinelAddresses) > 0 {
		client = rds.NewFailoverClient(&rds.FailoverOptions{
			MasterName:    cfg.MasterName,
			SentinelAddrs: cfg.SentinelAddresses,
			Password:      cfg.Password,
			DB:            cfg.Database,
			DialTimeout:   cfg.DialTimeout,
			ReadTimeout:   cfg.ReadTimeout,
			WriteTimeout:  cfg.WriteTimeout,
			PoolTimeout:   cfg.PoolTimeout,
			PoolSize:      cfg.RedisPoolSize,
			OnConnect: func(ctx context.Context, conn *rds.Conn) error {
				return conn.Ping(ctx).Err()
			},
		})
	} else {
		client = rds.NewClient(&rds.Options{
			Addr:         cfg.Address,
			Password:     cfg.Password,
			DB:           cfg.Database,
			DialTimeout:  cfg.DialTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			PoolTimeout:  cfg.PoolTimeout,
			PoolSize:     cfg.RedisPoolSize,
			OnConnect: func(ctx context.Context, conn *rds.Conn) error {
				return conn.Ping(ctx).Err()
			},
		})
	}

	err := client.Ping(context.Background()).Err()

	return client, err
}

// GetClient
func (rds *Redis) GetClient() *rds.Client {
	return rds.client
}

// Close pool of connections
func (rds *Redis) Close() error {
	return rds.client.Close()
}
