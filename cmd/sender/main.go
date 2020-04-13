package main

import (
	"github.com/mvanyushkin/go-calendar/config"
	"github.com/mvanyushkin/go-calendar/internal/sender"
	"github.com/mvanyushkin/go-calendar/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("The config file is broken: %v", err.Error())
	}

	logger.SetupLogger(cfg.LogFile, cfg.LogLevel)
	log.Info("application started.")

	s := sender.CreateSender(cfg.RabbitMQ)
	err = s.ListenMessages()
	if err != nil {
		log.Fatal(err)
	}
}
