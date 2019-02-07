package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"spx/utils"
)


func ProxyAuthenticateHandler(handle http.HandlerFunc, authUser, authPassword string) http.Handler {
	var proxyCredentials string
	var proxyCredentialsBase64 string

	if authUser != "" && authPassword != "" {
		proxyCredentials = fmt.Sprintf("%s:%s", authUser, authPassword)
	}
	if len(proxyCredentials) > 0 {
		proxyCredentialsBase64 = utils.Base64Encode(proxyCredentials)
	}

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		if len(proxyCredentials) == 0 {
			handle(w, r)
			return
		}

		// Need authorization
		credientials := strings.TrimPrefix(r.Header.Get("Proxy-Authorization"), "Basic ")
		if credientials == "" || (credientials != proxyCredentials && credientials != proxyCredentialsBase64){
			w.Header().Set("Proxy-Authenticate", "Basic realm=\"Password\"")
			http.Error(w, "", http.StatusProxyAuthRequired)
			log.Println("Accessing proxy deny, password is wrong or empty")
			return
		}
		handle(w, r)
	})
}
