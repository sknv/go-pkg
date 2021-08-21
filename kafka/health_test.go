package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingCheck(t *testing.T) {
	t.Run("returns an error for an empty producer", func(t *testing.T) {
		var prod SyncProducer
		check := prod.CheckHealth()
		assert.Error(t, check())
	})
}
