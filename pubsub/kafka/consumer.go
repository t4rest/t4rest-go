package kafka

import (
	"context"
	"sync"

	"github.com/t4rest/t4rest-go/logger"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

// Subscriber allows subscription.
type Subscriber struct {
	saramaConsumer sarama.ConsumerGroup
	handler        sarama.ConsumerGroupHandler
	cfg            Conf
	closeChan      chan struct{}
	consume        bool
	log            logger.Logger
}

// NewConsumer .
func NewConsumer(cfg Conf, handler sarama.ConsumerGroupHandler) *Subscriber {
	return &Subscriber{cfg: cfg, handler: handler}
}

// Start starts event module
func (sub *Subscriber) Start() error {
	sub.consume = true
	sub.closeChan = make(chan struct{})

	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	var err error
	sub.saramaConsumer, err = sarama.NewConsumerGroup(sub.cfg.Addresses, sub.cfg.ConsumerGroupID, config)
	if err != nil {
		return errors.Wrap(err, "Error creating consumer group client")
	}

	go func() {
		for err = range sub.saramaConsumer.Errors() {
			sub.log.Errorf("consume error: %s", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			if err := sub.saramaConsumer.Consume(ctx, sub.cfg.Topics, sub.handler); err != nil {
				sub.log.Fatalf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		sub.log.Debug("ctx.Done")
	case <-sub.closeChan:
		sub.log.Debug("sub.closeChan")
	}

	cancel()
	wg.Wait()

	return nil
}

// Stop stops event module
func (sub *Subscriber) Stop() error {

	sub.consume = false
	close(sub.closeChan)

	err := sub.saramaConsumer.Close()
	if err != nil {
		return errors.Wrap(err, "sub.saramaConsumer.Close")
	}

	return nil
}

// Title returns events title.
func (sub *Subscriber) Title() string {
	return "Subscriber"
}
