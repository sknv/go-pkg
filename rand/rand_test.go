package rand

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	gofakeit.Seed(0)

	t.Run("generates a string with proper length", func(t *testing.T) {
		length := int(gofakeit.Int8())
		str := String(length)
		assert.Equal(t, length, len(str))
	})
}
