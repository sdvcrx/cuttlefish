package main

import (
	"log"
	"time"
	"net"
	"net/url"
	"net/http"
	"spx/config"
)

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

func selectProxy(r *http.Request) (*url.URL, error) {
	proxy := config.GetInstance().ParentProxies.Next()
	if proxy != "" {
		log.Printf("select proxy: %s", proxy)
		// return url.Parse(proxy)
	}
	return nil, nil
}
