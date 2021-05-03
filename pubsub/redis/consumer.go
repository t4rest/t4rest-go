package redis

import (
	"context"

	rds "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/t4rest/t4rest-go/redis"
)

// Subscriber allows subscription.
type Subscriber struct {
	rds    *rds.Client
	cfg    redis.Conf
	pubsub *rds.PubSub
	ch     <-chan *rds.Message
	topic  string
	handle Handler
}

type Handler func(msg *rds.Message)

// New return new subscriber
func New(cfg redis.Conf, handle Handler) (*Subscriber, error) {

	r, err := redis.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "redis.NewClient")
	}

	return &Subscriber{cfg: cfg, rds: r, handle: handle}, nil
}

// Title returns events title.
func (sub *Subscriber) Title() string {
	return "Subscriber Rabbit"
}

// Start starts event module
func (sub *Subscriber) Start() error {

	sub.pubsub = sub.rds.Subscribe(context.Background(), "mychannel1")

	// Go channel which receives messages.
	sub.ch = sub.pubsub.Channel()
	for msg := range sub.ch {
		sub.handle(msg)
	}

	return nil
}

// Stop stops event module
func (sub *Subscriber) Stop() error {
	return sub.pubsub.Close()
}
