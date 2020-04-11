package logger

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func SetupLogger(filePath string, logLevel string) {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log.SetOutput(os.Stdout)
	if err == nil {
		log.SetOutput(io.MultiWriter(os.Stdout, f))
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Infof("Unknown log level %v", logLevel)
	}
	log.SetLevel(level)
}
