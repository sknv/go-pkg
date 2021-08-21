package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"

	"github.com/sknv/go-pkg/closer"
)

// SyncProducer is a helper wrapper for sarama.SyncProducer using sarama.ByteEncoder.
type SyncProducer struct {
	sarama.SyncProducer

	client sarama.Client
}

// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
func NewSyncProducer(brokers []string, config Config) (*SyncProducer, error) {
	client, err := sarama.NewClient(brokers, config.Sarama)
	if err != nil {
		return nil, fmt.Errorf("create sarama client: %w", err)
	}

	prod, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, fmt.Errorf("create sarama sync producer from client: %w", err)
	}

	// Register tracing if needed
	if config.EnableTracing {
		prod = otelsarama.WrapSyncProducer(config.Sarama, prod)
	}

	return &SyncProducer{
		SyncProducer: prod,

		client: client,
	}, nil
}

// Publish publishes a message to the provided topic.
func (p *SyncProducer) Publish(
	_ context.Context, message *sarama.ProducerMessage,
) (partition int32, offset int64, err error) {
	return p.SendMessage(message)
}

// IsHealthy signals the producer is healthy.
func (p *SyncProducer) IsHealthy() bool {
	if p.client == nil {
		return false
	}
	return len(p.client.Brokers()) > 0
}

// Close tries to close the producer gracefully.
func (p *SyncProducer) Close(ctx context.Context) error {
	if err := closer.CloseContext(ctx, p.SyncProducer.Close); err != nil {
		return fmt.Errorf("close sync producer: %w", err)
	}
	return nil
}
