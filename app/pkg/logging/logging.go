package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	Logger *zerolog.Logger
}

func (logger *Logger) Ftl(msg string) {
	logger.Logger.Fatal().
		CallerSkipFrame(3).Msg(msg)
}

func (logger *Logger) Request(method, path string, status int, latency time.Duration) {
	logger.Logger.Info().
		Str("method", method).
		Int("status", status).
		Str("path", path).
		Dur("latency", latency).
		Msg("Request")
}

func (logger *Logger) Err(errKey string, errMsg string, fileInfo string) {
	if fileInfo == "" {
		logger.Logger.Error().Msgf("%s: %s", errKey, errMsg)
	} else {
		logger.Logger.Error().Str("cd", fileInfo).Msgf("%s: %s", errKey, errMsg)
	}
}

func (logger *Logger) Inf(msg string) {
	logger.Logger.Info().Msg(msg)
}

var defaultLogger = newLogger()

func newLogger() *zerolog.Logger {
	output := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.Kitchen
		w.Out = os.Stdout
		w.FormatLevel = func(i interface{}) string {
			level := strings.ToUpper(fmt.Sprintf("[%s]", i))
			switch i {
			case "debug":
				return "\x1b[35m" + level + "\x1b[37m"
			case "info":
				return "\x1b[32m" + level + "\x1b[37m"
			case "warn":
				return "\x1b[33m" + level + "\x1b[37m"
			case "error":
				return "\x1b[31m" + level + "\x1b[37m"
			case "fatal":
				return "\x1b[31m" + level + "\x1b[37m"
			default:
				return "\x1b[0m" + level
			}
		}
		w.FormatMessage = func(i interface{}) string {
			message := fmt.Sprintf("--- %v ---", i)
			return "\x1b[32m" + message + "\x1b[0m"
		}
	})

	_, fileName, line, _ := runtime.Caller(0)

	err := os.MkdirAll("app/pkg/storage/logs", 0755)
	if err != nil {
		log.Fatalf("Error occured on find directory: %v %v", fileName, line)
	}

	file, err := os.OpenFile("app/pkg/storage/logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatalf("Error occured on file open/create: %v %v", fileName, line)
	}

	multi := io.MultiWriter(zerolog.ConsoleWriter{Out: file}, output)
	logger := zerolog.New(multi).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
	return &logger
}

func GetLogger() *Logger {
	return &Logger{defaultLogger}
}
