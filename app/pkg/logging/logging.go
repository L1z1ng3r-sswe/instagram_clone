package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	Logger *zerolog.Logger
}

func (logger *Logger) Ftl(msg string) {
	logger.Logger.Fatal().Msg(msg)
}

var defaultLogger = newLogger()

func newLogger() *zerolog.Logger {
	output := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.Kitchen
		w.Out = os.Stdout
		w.FormatMessage = func(i interface{}) string {
			message := fmt.Sprintf("--- %v ---", i)
			return "\x1b[100m" + message + "\x1b[0m"
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
