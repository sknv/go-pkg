package rand

import (
	"math/rand"
	"time"
)

const (
	_letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_digits  = "0123456789"
	_charset = _digits + _letters
)

//nolint:gosec // no need in secure random
var _seededRand = rand.New(
	rand.NewSource(
		time.Now().UnixNano(),
	),
)

// String generates a random string with a provided length.
func String(length int) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = _charset[_seededRand.Int()%len(_charset)]
	}
	return string(buf)
}
