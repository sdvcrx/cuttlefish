package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sdvcrx/cuttlefish/config"
	"github.com/sdvcrx/cuttlefish/log"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	logger  = log.Logger
)

func main() {
	if viper.GetBool("version") {
		fmt.Printf("%s %s %s\n", version, commit, date)
		return
	}
	log.SetLevel(viper.GetBool("verbose"))

	config.Load()

	InitSignals()

	server := NewProxyServer()
	logger.Info().Msgf("Proxy server is listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal().Err(errors.Wrap(err, "server.ListenAndServe")).Msg("Failed to start proxy server")
	}
}
