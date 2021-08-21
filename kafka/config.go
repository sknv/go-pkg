package kafka

import (
	"github.com/Shopify/sarama"
)

// NewDefaultConfig creates a new default config.
func NewDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	return config
}
