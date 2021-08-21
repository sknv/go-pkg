package kafka

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/korovkin/limiter"

	"github.com/sknv/go-pkg/log"
)

// ConsumerHandler handles a message from Kafka.
type ConsumerHandler interface {
	// Consume consumes a *sarama.Message. If an error is returned the message will not be marked as consumed.
	Consume(context.Context, *sarama.ConsumerMessage) error
	// Limit sets the concurrency limit for the handler. Non-positive value means the default level = 100.
	Limit() int
}

type consumerHandler struct {
	handler ConsumerHandler
	ready   chan struct{}
}

func newConsumerHandler(handler ConsumerHandler) *consumerHandler {
	return &consumerHandler{
		handler: handler,
		ready:   make(chan struct{}),
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready) // mark the consumer as ready
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (c *consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing loop and exit.
func (c *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := session.Context()

	limit := limiter.NewConcurrencyLimiter(c.handler.Limit())
	defer limit.Wait()

	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		limit.Execute(func() {
			if err := c.handler.Consume(ctx, message); err != nil {
				log.Extract(ctx).WithError(err).Error("handle kafka message")
				return
			}
			session.MarkMessage(message, "")
		})
	}
	return nil
}

func (c *consumerHandler) reset() {
	c.ready = make(chan struct{})
}

func (c *consumerHandler) waitUntilReady() {
	<-c.ready
}
