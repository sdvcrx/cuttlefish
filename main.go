package main

import (
	"fmt"
	"log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"spx/config"
)

var (
	version = "dev"
	commit = "none"
	date = "unknown"
)

func main() {
	config.Load()

	if viper.GetBool("version") {
		fmt.Printf("%s %s %s\n", version, commit, date)
		return
	}

	server := NewProxyServer()
	log.Printf("Proxy server is listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(errors.Wrap(err, "server.ListenAndServe"))
	}
}
