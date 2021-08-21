package kafka

import (
	"errors"

	"github.com/heptiolabs/healthcheck"
)

// CheckHealth returns a healthcheck.Check that validates connectivity to a kafka client.
func (p *SyncProducer) CheckHealth() healthcheck.Check {
	return func() error {
		if !p.IsHealthy() {
			return errors.New("kafka is down")
		}
		return nil
	}
}
