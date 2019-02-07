package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleProxyIterator(t *testing.T) {
	proxies := []string{"a", "b", "c"}
	iter := NewProxyIterator(proxies)
	for i := 0; i < 4; i++ {
		assert.Contains(t, proxies, iter.Next())
	}
}

func TestEmptyProxyIterator(t *testing.T) {
	iter := NewProxyIterator([]string{})
	for i := 0; i < 4; i++ {
		assert.Empty(t, iter.Next())
	}
}
