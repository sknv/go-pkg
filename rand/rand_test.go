package rand

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("generates a string with proper length", func(t *testing.T) {
		length := rand.Intn(16) //nolint:gosec // no need in secure random
		str := String(length)
		assert.Equal(t, length, len(str))
	})
}
