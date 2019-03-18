package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/sdvcrx/cuttlefish/config"
)

func InitSignals() {
	// handle signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP)

	go func() {
		for _ = range signalChan {
			config.Reload()
		}
	}()
}
