package main

import (
	"time"
	"net"
	"net/url"
	"net/http"
)

func selectProxy(r *http.Request) (*url.URL, error) {
	// return url.Parse("http://127.0.0.1:7890")
	return nil, nil
}

var DefaultProxyTransport = &http.Transport{
	Proxy: selectProxy,
	// Copy from https://golang.org/pkg/net/http/#RoundTripper
	// http.DefaultTransport
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
