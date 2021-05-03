package redis

import (
	"context"

	rds "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/t4rest/t4rest-go/redis"
)

type redisPub struct {
	rds    *redis.Redis
	cfg    redis.Conf
	pubsub *rds.PubSub
}

// NewProducer .
func NewProducer(cfg redis.Conf) (*redisPub, error) {

	rdsPub := &redisPub{cfg: cfg}

	rc, err := redis.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "redis.New")
	}

	rdsPub.pubsub = rc.Subscribe(context.Background(), "mychannel1")

	return rdsPub, nil
}

// Stop stops event module
func (pub *redisPub) Stop() error {
	return pub.pubsub.Close()
}
