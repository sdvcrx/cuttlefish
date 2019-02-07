package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"spx/utils"
	"testing"
)

func assertHTTPCode(t *testing.T, handler http.HandlerFunc, username, password string, expectedCode int) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	if username != "" && password != "" {
		req.Header.Set("Proxy-Authorization", "Basic "+utils.Base64Encode(username+":"+password))
	}
	handler(w, req)

	assert.Equal(t, expectedCode, w.Code)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func TestProxyAuthenicateHandlerWithoutAuth(t *testing.T) {
	px := ProxyAuthenticateHandler(handlerFunc, "", "")
	assert.HTTPBodyContains(t, px.ServeHTTP, "GET", "/", nil, "ok")
}

func TestProxyAuthenicateHandler(t *testing.T) {
	px := ProxyAuthenticateHandler(handlerFunc, "test", "test")
	assertHTTPCode(t, px.ServeHTTP, "test", "test", http.StatusOK)
}

func TestProxyAuthenicateHandlerDeny(t *testing.T) {
	px := ProxyAuthenticateHandler(handlerFunc, "fail", "fail")
	assertHTTPCode(t, px.ServeHTTP, "joida", "asdwe", http.StatusProxyAuthRequired)
}
