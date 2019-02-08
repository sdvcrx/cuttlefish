package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sdvcrx/cuttlefish/config"
	"github.com/spf13/viper"
	"log"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if viper.GetBool("version") {
		fmt.Printf("%s %s %s\n", version, commit, date)
		return
	}

	config.Load()

	server := NewProxyServer()
	log.Printf("Proxy server is listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(errors.Wrap(err, "server.ListenAndServe"))
	}
}
