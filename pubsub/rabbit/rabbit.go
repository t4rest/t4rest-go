package rabbit

import (
	"context"

	"github.com/streadway/amqp"
)

// Handler .
type Handler func(msg amqp.Delivery)

// Conf .
type Conf struct {
	ConnectionString      string
	Username              string
	Password              string
	ConsumeQueueName      string
	ConsumerWorkers       int
	ConsumerPrefetchCount int
}

// PublishOptions .
type PublishOptions struct {
	WorkerPubRouteKey string
	ProduceExchange   string
	CorrelationID     string
}

// Publisher interface
type Publisher interface {
	Publish(ctx context.Context, data []byte, opt PublishOptions) error
	Close() error
}
