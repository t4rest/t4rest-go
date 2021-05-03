package kafka_confluent

import (
	"strings"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/t4rest/t4rest-go/logger"
)

// Subscriber allows subscription.
type Subscriber struct {
	consumer  *kafka.Consumer
	cfg       Conf
	connected bool
	closeCh   chan bool
	mx        sync.Mutex
	kafkaCfg  *kafka.ConfigMap
	log       *logger.Logger
	handler   Handler
}

// NewConsumer return new subscriber
func NewConsumer(cfg Conf, handler Handler, log *logger.Logger) *Subscriber {
	return &Subscriber{cfg: cfg, handler: handler, log: log}
}

// Title .
func (sub *Subscriber) Title() string {
	return "Subscriber"
}

// Start event module
func (sub *Subscriber) Start() error {

	chLogEvent := make(chan kafka.LogEvent)
	sub.kafkaCfg = &kafka.ConfigMap{
		"bootstrap.servers":      strings.Join(sub.cfg.Addresses, ","),
		"client.id":              sub.cfg.AppID,
		"session.timeout.ms":     sub.cfg.SessionTimeoutMs,    // Group.Session.Timeout
		"heartbeat.interval.ms":  sub.cfg.HeartbeatIntervalMs, // Group.Heartbeat.Interval
		"socket.timeout.ms":      sub.cfg.SocketTimeoutMs,     // Net.DialTimeout Net.ReadTimeout
		"group.id":               sub.cfg.ConsumerGroupID,
		"auto.offset.reset":      "latest",
		"enable.partition.eof":   true,
		"enable.auto.commit":     false,
		"go.logs.channel.enable": true,
		"go.logs.channel":        chLogEvent,
	}

	go func() {
		for le := range chLogEvent {
			sub.log.With("Message", le.Message, "Tag", le.Tag, "Name", le.Name, "Timestamp", le.Timestamp).
				Error("consumer.Logs")
		}
	}()

	return sub.consume()

}

func (sub *Subscriber) consume() error {
	var err error

	sub.consumer, err = kafka.NewConsumer(sub.kafkaCfg)
	if err != nil {
		sub.log.With("error", err).Fatal("NewConsumer")
	}

	err = sub.consumer.Subscribe(sub.cfg.Topic, nil)
	if err != nil {
		sub.log.With("error", err).Fatal("Subscribe")
	}

	for sub.connected {
		ev := sub.consumer.Poll(100)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			sub.handler(e)
		case kafka.PartitionEOF:
			sub.log.Debug("Kafka Poll PartitionEOF")
		case kafka.Error:
			sub.log.With("code", e.Code(), "err", e.String()).Error("kafka poll error event")
		}
	}

	err = sub.consumer.Close()
	close(sub.closeCh)

	return err
}

// Stop event module
func (sub *Subscriber) Stop() error {
	err := sub.consumer.Unsubscribe()
	sub.connected = false

	// wait for messages
	<-sub.closeCh

	return err
}
