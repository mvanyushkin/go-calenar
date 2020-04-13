package main

import (
	"github.com/mvanyushkin/go-calendar/config"
	"github.com/mvanyushkin/go-calendar/internal/reminder"
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

	reminder := reminder.CreateReminder(cfg.ConnectionString, "amqp://user:aA123456@localhost:5672/")
	err = reminder.Do()
	if err != nil {
		log.Fatal(err)
	}
}
