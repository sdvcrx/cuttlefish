package main

import (
	"fmt"
	"github.com/sdvcrx/cuttlefish/utils"
	"net/http"
	"strings"
)

type BasicAuth struct {
	username          string
	password          string
	credentials       string
	credentialsBase64 string
}

func (auth BasicAuth) IsEmpty() bool {
	return auth.username == "" && auth.password == ""
}

func (auth BasicAuth) Validate(authHeader string) bool {
	if auth.IsEmpty() {
		return true
	}
	credientials := strings.TrimPrefix(authHeader, "Basic ")
	if credientials == auth.credentials || credientials == auth.credentialsBase64 {
		return true
	}
	return false
}

func NewBasicAuth(username, password string) BasicAuth {
	var proxyCredentials string
	var proxyCredentialsBase64 string

	if username != "" && password != "" {
		proxyCredentials = fmt.Sprintf("%s:%s", username, password)
	}
	if proxyCredentials != "" {
		proxyCredentialsBase64 = utils.Base64Encode(proxyCredentials)
	}
	return BasicAuth{
		username,
		password,
		proxyCredentials,
		proxyCredentialsBase64,
	}
}

func ProxyAuthenticateHandler(handle http.HandlerFunc, authUser, authPassword string) http.Handler {
	auth := NewBasicAuth(authUser, authPassword)

	if auth.IsEmpty() {
		return handle
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Need authorization
		if !auth.Validate(r.Header.Get("Proxy-Authorization")) {
			w.Header().Set("Proxy-Authenticate", "Basic realm=\"Password\"")
			http.Error(w, "", http.StatusProxyAuthRequired)
			logger.Error().Msgf("Accessing proxy deny, password is wrong or empty")
			return
		}
		handle(w, r)
	})
}
