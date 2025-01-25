package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
)

func New(level string) *logrus.Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	switch level {
	case DEBUG:
		logger.SetLevel(logrus.DebugLevel)
	case INFO:
		logger.SetLevel(logrus.InfoLevel)
	default:
		log.Fatal("Cannot set log level")
	}

	return logger
}
