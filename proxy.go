package main

import (
	"bufio"
	"fmt"
	"github.com/sdvcrx/cuttlefish/config"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	FILTER_HEADERS = []string{
		"Prxoy-Authenticate",
		"Proxy-Connection",
		"Transfer-Encoding",
		"Upgrade",
		"X-Forwarded-For",
	}
)

func transfer(dst io.WriteCloser, src io.ReadCloser) {
	defer dst.Close()
	defer src.Close()
	io.Copy(dst, src)
}

func copyHeader(dst, src http.Header) {
	for key, vv := range src {
		for _, val := range vv {
			dst.Add(key, val)
		}
	}
}

func filterProxyHeaders(header http.Header) {
	for _, key := range FILTER_HEADERS {
		header.Del(key)
	}
}

func connectTunnelHandler(w http.ResponseWriter, r *http.Request) {
	parentProxy, err := SelectProxy(r)

	var proxy_conn net.Conn

	if parentProxy != nil {
		proxy_conn, err = net.DialTimeout("tcp4", parentProxy.Host, 10*time.Second)
		// TODO remove current unavailable proxy from proxies pool
	} else {
		proxy_conn, err = net.DialTimeout("tcp4", r.Host, 10*time.Second)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// if parentProxy is not null, send CONNECT request to parent proxy server
	if parentProxy != nil {
		connectReq := &http.Request{
			Method: http.MethodConnect,
			URL:    r.URL,
			Host:   r.Host,
			Header: make(http.Header),
		}
		connectReq.Write(proxy_conn)
		br := bufio.NewReader(proxy_conn)
		resp, err := http.ReadResponse(br, connectReq)
		if err != nil {
			proxy_conn.Close()
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		// expect 200 Connection established
		if resp.StatusCode != http.StatusOK {
			proxy_conn.Close()
			// TODO remove current unavailable proxy from proxies pool
			http.Error(w, "Read Proxy Response error", http.StatusServiceUnavailable)
			return
		}
	}

	// Send 200 Connection established
	w.WriteHeader(http.StatusOK)

	// hijack connnection
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Proxy hijack error", http.StatusServiceUnavailable)
		return
	}
	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	go transfer(proxy_conn, client_conn)
	go transfer(client_conn, proxy_conn)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	filterProxyHeaders(r.Header)
	proxy_resp, err := DefaultProxyTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer proxy_resp.Body.Close()

	copyHeader(w.Header(), proxy_resp.Header)
	w.WriteHeader(proxy_resp.StatusCode)
	io.Copy(w, proxy_resp.Body)
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		log.Printf("%s %s", r.Method, r.Host)
		connectTunnelHandler(w, r)
	} else {
		log.Printf("%s %s", r.Method, r.URL)
		httpHandler(w, r)
	}
}

func NewProxyServer() http.Server {
	appConfig := config.GetInstance()
	authUser := appConfig.AuthUser
	authPassword := appConfig.AuthPassword

	addr := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	proxyHandler := ProxyAuthenticateHandler(ProxyHandler, authUser, authPassword)
	return http.Server{
		Addr:    addr,
		Handler: proxyHandler,
	}
}
