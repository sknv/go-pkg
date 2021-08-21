package kafka

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"

	"github.com/sknv/go-pkg/log"
)

// ConsumerGroup is a helper wrapper for sarama.ConsumerGroup.
type ConsumerGroup struct {
	sarama.ConsumerGroup

	enableTracing bool
	cancel        context.CancelFunc
}

// NewConsumerGroup returns a new instance.
func NewConsumerGroup(brokers []string, group string, config Config) (*ConsumerGroup, error) {
	cons, err := sarama.NewConsumerGroup(brokers, group, config.Sarama)
	if err != nil {
		return nil, err
	}

	return &ConsumerGroup{
		ConsumerGroup: cons,

		enableTracing: config.EnableTracing,
	}, nil
}

// Consume starts consuming from the provided topics.
func (c *ConsumerGroup) Consume(ctx context.Context, topics []string, handler ConsumerHandler, wg *sync.WaitGroup) {
	wg.Add(1)

	ctx, c.cancel = context.WithCancel(ctx)
	cons := newConsumerHandler(handler)

	// Register tracing if needed
	var consumerHandler sarama.ConsumerGroupHandler = cons
	if c.enableTracing {
		consumerHandler = otelsarama.WrapConsumerGroupHandler(consumerHandler)
	}

	go func() {
		defer wg.Done()

		logger := log.Extract(ctx).WithField("topics", topics)

		// `Consume` should be called inside an infinite loop, when a server-side rebalance happens,
		// the consumer session will need to be recreated to get the new claims
		for {
			if err := c.ConsumerGroup.Consume(ctx, topics, consumerHandler); err != nil {
				logger.WithError(err).Error("consume from topics")
			}

			// Check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				logger.Info("consumption stopped")
				return
			}

			cons.reset()
		}
	}()

	cons.waitUntilReady()
}

// Close closes the consumer.
func (c *ConsumerGroup) Close() error {
	c.cancel()
	return c.ConsumerGroup.Close()
}
