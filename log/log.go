package log

import (
	"github.com/rs/zerolog"
	"os"
)

var Logger zerolog.Logger

func SetLevel(verbose bool) {
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func init() {
	out := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	Logger = zerolog.New(out).With().Timestamp().Logger()
}
