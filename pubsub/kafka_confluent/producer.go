package kafka_confluent

import (
	"context"
	"strings"
	"time"

	"github.com/t4rest/t4rest-go/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// kafkaPub kafka publisher
type kafkaPub struct {
	producer *kafka.Producer
	cfg      Conf
	log      *logger.Logger
}

// NewProducer creates new kafka connection
func NewProducer(cfg Conf, log *logger.Logger) (Publisher, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  strings.Join(cfg.Addresses, ","),
		"client.id":          cfg.AppID,
		"request.timeout.ms": cfg.RequestTimeoutMs,
		"socket.timeout.ms":  cfg.SocketTimeoutMs, // Net.DialTimeout Net.ReadTimeout Net.WriteTimeout
	})

	if err != nil {
		return nil, err
	}

	return &kafkaPub{producer: p, cfg: cfg, log: log}, nil
}

// Publish .
func (kfk *kafkaPub) Publish(ctx context.Context, data []byte, opt PublishOptions) error {

	deliveryChan := make(chan kafka.Event)

	err := kfk.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &opt.Topic, Partition: kafka.PartitionAny},
		Value:          data,
		Key:            []byte(opt.Key),
		Timestamp:      time.Now().UTC(),
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		kfk.log.Errorf("Delivery failed: %v", m.TopicPartition.Error)
	} else {
		kfk.log.Debugf("Delivered message to topic %s [%d] at offset %v",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)

	return err
}

// Close connection
func (kfk *kafkaPub) Close() error {
	kfk.producer.Flush(int(time.Second))
	kfk.producer.Close()
	return nil
}
