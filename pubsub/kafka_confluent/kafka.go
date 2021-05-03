package kafka_confluent

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Handler .
type Handler func(msg *kafka.Message)

// Publisher interface
type Publisher interface {
	Publish(ctx context.Context, data []byte, opt PublishOptions) error
	Close() error
}

// PublishOptions .
type PublishOptions struct {
	Topic string
	Key   string
}

// Conf .
type Conf struct {
	Addresses           []string
	Topics              []string
	AppID               string
	ConsumerGroupID     string
	Topic               string
	RequestTimeoutMs    int
	SessionTimeoutMs    int
	SocketTimeoutMs     int
	HeartbeatIntervalMs int
}
