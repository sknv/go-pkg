package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterTracer(t *testing.T) {
	t.Run("registers tracing successfully", func(t *testing.T) {
		_, err := RegisterTracer(DriverName)
		assert.NoError(t, err)
	})
}
