package rabbit

import (
	"fmt"

	"github.com/t4rest/t4rest-go/logger"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Subscriber allows subscription.
type Subscriber struct {
	cfg       Conf
	closeChan chan struct{}
	conn      *amqp.Connection
	ch        *amqp.Channel
	log       *logger.Logger
	handle    Handler
	isClosed  bool
}

// New return new subscriber
func New(cfg Conf, log *logger.Logger) *Subscriber {
	return &Subscriber{
		cfg: cfg,
		log: log,
	}
}

// Title returns events title.
func (sub *Subscriber) Title() string {
	return "Subscriber Rabbit"
}

// Start starts event module
func (sub *Subscriber) Start() error {
	sub.closeChan = make(chan struct{})

	var err error
	sub.conn, err = amqp.Dial(fmt.Sprintf(sub.cfg.ConnectionString, sub.cfg.Username, sub.cfg.Password))
	if err != nil {
		return errors.Wrap(err, "ampq.dial")
	}

	sub.ch, err = sub.conn.Channel()
	if err != nil {
		return errors.Wrap(err, "ampq.chanel")
	}

	errChan := make(chan *amqp.Error)
	sub.ch.NotifyClose(errChan)
	go func() {
		err = <-errChan
		close(sub.closeChan)

		sub.log.With("component", "ampq.chanel").Debug("NotifyClose")
	}()

	err = sub.ch.Qos(sub.cfg.ConsumerPrefetchCount, 0, false)
	if err != nil {
		return errors.Wrap(err, "ampq.qos")
	}

	for i := 0; i < sub.cfg.ConsumerWorkers; i++ {
		go func() {
			msgs, err := sub.ch.Consume(
				sub.cfg.ConsumeQueueName, // queue
				"",                       // consumer
				false,                    // auto-ack
				false,                    // exclusive
				false,                    // no-local
				false,                    // no-wait
				nil,                      // args
			)
			if err != nil {
				sub.log.With(err).Error("ampq.consume")
				return
			}

			for m := range msgs {
				sub.handle(m)
			}
		}()
	}

	<-sub.closeChan

	return nil
}

// Stop stops event module
func (sub *Subscriber) Stop() error {
	sub.isClosed = true

	if sub.ch != nil {
		err := sub.ch.Close()
		if err != nil {
			return errors.Wrap(err, "sub.ch.Close")
		}
	}

	if sub.conn != nil {
		err := sub.conn.Close()
		if err != nil {
			return errors.Wrap(err, "sub.conn.Close")
		}
	}
	return nil
}
