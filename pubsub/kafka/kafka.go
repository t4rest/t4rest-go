package kafka

import "context"

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
	ConsumerGroupID     string
	RequestTimeoutMs    int
	SessionTimeoutMs    int
	SocketTimeoutMs     int
	HeartbeatIntervalMs int
}
