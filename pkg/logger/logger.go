package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func NewDefaultConsoleLogger() *zerolog.Logger {
	consoleOutput := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
		w.Out = os.Stdout
	})
	logger := zerolog.New(consoleOutput)
	logger = logger.With().Timestamp().Logger()
	return &logger
}
