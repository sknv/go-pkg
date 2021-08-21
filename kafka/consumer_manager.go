package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/sknv/go-pkg/closer"
	"github.com/sknv/go-pkg/log"
)

// ConsumerManager helps to manage multiple consumers.
type ConsumerManager struct {
	closers *closer.Closers
	wg      sync.WaitGroup
}

// NewConsumerManager returns a new instance.
func NewConsumerManager() *ConsumerManager {
	return &ConsumerManager{
		closers: closer.NewClosers(),
	}
}

// RegisterConsumerGroup registers new consumer group and returns it.
func (c *ConsumerManager) RegisterConsumerGroup(brokers []string, group string, config Config) (*ConsumerGroup, error) {
	cons, err := NewConsumerGroup(brokers, group, config)
	if err != nil {
		return nil, err
	}

	c.closers.Add(func(ctx context.Context) error {
		logger := log.Extract(ctx).WithField("group", group)
		logger.Info("stopping kafka consumer...")
		defer logger.Info("kafka consumer stopped")

		return cons.Close()
	})

	return cons, nil
}

// Consume starts consuming from the provided topics.
func (c *ConsumerManager) Consume(
	ctx context.Context, consumer *ConsumerGroup, topics []string, handler ConsumerHandler,
) {
	consumer.Consume(ctx, topics, handler, &c.wg)
}

// Close tries to close the consumers gracefully.
func (c *ConsumerManager) Close(ctx context.Context) error {
	var err error
	closed := make(chan struct{})
	go func() {
		err = c.closers.Close(ctx)
		c.wg.Wait()
		close(closed)
	}()

	select {
	case <-closed:
		return err
	case <-ctx.Done():
		return fmt.Errorf("close consumers: %w", ctx.Err())
	}
}
