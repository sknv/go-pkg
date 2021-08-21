package kafka

import (
	"github.com/Shopify/sarama"
)

type Config struct {
	Sarama        *sarama.Config
	EnableTracing bool
}

// NewDefaultConfig creates a new default config.
func NewDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	return config
}
