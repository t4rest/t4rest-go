package kafka

import (
	"context"
	"time"

	"github.com/Shopify/sarama"
)

type kafkaPub struct {
	producer sarama.SyncProducer
}

// NewProducer .
func NewProducer(cfg Conf) (Publisher, error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(cfg.Addresses, config)
	if err != nil {
		return nil, err
	}

	return &kafkaPub{producer: producer}, nil
}

// Publish publish event
func (kfk *kafkaPub) Publish(_ context.Context, data []byte, opt PublishOptions) error {

	_, _, err := kfk.producer.SendMessage(&sarama.ProducerMessage{
		Topic:     opt.Topic,
		Value:     sarama.ByteEncoder(data),
		Key:       sarama.StringEncoder(opt.Key),
		Timestamp: time.Now(),
	})

	return err
}

// Close .
func (kfk *kafkaPub) Close() error {
	return kfk.producer.Close()
}
