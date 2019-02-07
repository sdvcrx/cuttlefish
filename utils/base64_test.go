package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	assert.Equal(t, Base64Encode("test:test"), "dGVzdDp0ZXN0")
}

func TestBase64Decode(t *testing.T) {
	res, err := Base64Decode("dGVzdDp0ZXN0")
	assert.Nil(t, err)
	assert.Equal(t, res, "test:test")
}
