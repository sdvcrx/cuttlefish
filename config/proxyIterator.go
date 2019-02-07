package config

import (
	"math/rand"
	"time"
)

type Iterator interface {
	Next() string
}

// Random proxies iterator
type ProxyIterator struct {
	data []string
	len  int
	rd   *rand.Rand
}

func (iter ProxyIterator) Next() string {
	if iter.len == 0 {
		return ""
	}
	return iter.data[iter.rd.Intn(iter.len)]
}

func NewProxyIterator(proxies []string) ProxyIterator {
	rd := rand.New(rand.NewSource(time.Now().Unix()))

	return ProxyIterator{
		proxies,
		len(proxies),
		rd,
	}
}
