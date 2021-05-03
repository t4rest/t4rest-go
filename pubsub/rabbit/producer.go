package rabbit

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type rabbitPub struct {
	conn        *amqp.Connection
	ch          *amqp.Channel
	cfg         Conf
	isConnected bool
}

// NewProducer .
func NewProducer(cfg Conf) (Publisher, error) {
	rabbit := &rabbitPub{cfg: cfg}
	err := rabbit.connect()
	if err != nil {
		return nil, err
	}

	return rabbit, nil
}

// Publish .
func (pub *rabbitPub) Publish(_ context.Context, data []byte, opt PublishOptions) error {
	err := pub.maybeReconnect()
	if err != nil {
		return errors.Wrap(err, "pub.maybeReconnect")
	}

	return pub.ch.Publish(
		opt.ProduceExchange,   // exchange
		opt.WorkerPubRouteKey, // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			CorrelationId: opt.CorrelationID,
			Timestamp:     time.Now(),
			DeliveryMode:  amqp.Persistent,
			ContentType:   "application/json",
			Body:          data,
		})
}

// Close close connection
func (pub *rabbitPub) Close() error {
	if pub.ch != nil {
		err := pub.ch.Close()
		if err != nil {
			return errors.Wrap(err, "pub.ch.Close")
		}
	}

	if pub.conn != nil {
		err := pub.conn.Close()
		if err != nil {
			return errors.Wrap(err, "pub.conn.Close")
		}
	}

	return nil
}

func (pub *rabbitPub) connect() error {
	var err error

	pub.conn, err = amqp.Dial(fmt.Sprintf(pub.cfg.ConnectionString, pub.cfg.Username, pub.cfg.Password))
	if err != nil {
		return errors.Wrap(err, "amqp.Dial")
	}

	pub.ch, err = pub.conn.Channel()
	if err != nil {
		errClose := pub.Close()
		return errors.Wrapf(err, "pub.conn.Channel. pub.Close err: %s", errClose)
	}

	errChan := make(chan *amqp.Error)
	pub.ch.NotifyClose(errChan)

	pub.isConnected = true

	go func() {
		<-errChan
		pub.isConnected = false
	}()

	return nil
}

func (pub *rabbitPub) maybeReconnect() error {
	if pub.isConnected {
		return nil
	}

	return pub.connect()
}
