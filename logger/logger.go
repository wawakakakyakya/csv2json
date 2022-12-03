package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	process string
	logger  *zerolog.Logger
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func NewLogger(process string) *Logger {
	var logger zerolog.Logger
	writer := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false})
	logger = zerolog.New(writer).Level(zerolog.InfoLevel).With().
		Timestamp().
		Str("process", process).
		Logger()
	return &Logger{process: process, logger: &logger}
}
