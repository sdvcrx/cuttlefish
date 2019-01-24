package main

import (
	"log"
	"github.com/pkg/errors"
	"spx/config"
)

func main() {
	config.Load()

	server := NewProxyServer()
	log.Printf("Proxy server is listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(errors.Wrap(err, "server.ListenAndServe"))
	}
}
