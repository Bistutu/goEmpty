package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncry(t *testing.T) {
	key := "k0SlH7X2OGNBpmVT"
	msg := "123456"
	en, err := Encrypt(msg, key)
	assert.NoError(t, err)
	de := Decrypt(en, key)
	assert.Equal(t, msg, de)
}

func TestMyEncry(t *testing.T) {
	key := "D7NUPv3EozQiZH6Q"
	msg := "aaa212265"
	en, err := Encrypt(msg, key)
	assert.NoError(t, err)
	de := Decrypt(en, key)
	assert.Equal(t, msg, de)
}

func BenchmarkEncry(b *testing.B) {
	key := "k0SlH7X2OGNBpmVT"
	msg := "123456"
	for i := 0; i < b.N; i++ {
		en, err := Encrypt(msg, key)
		assert.NoError(b, err)
		assert.NotEmpty(b, en)
	}
}
