package main

import (
	"context"
	"flag"
	"github.com/mvanyushkin/go-calendar/internal/config"
	"github.com/mvanyushkin/go-calendar/internal/reminder"
	"github.com/mvanyushkin/go-calendar/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configFilePath := flag.String("config", "", "settings file")
	flag.Parse()
	if configFilePath == nil {
		defaultConfigFileName := "local_config.json"
		configFilePath = &defaultConfigFileName
	}

	cfg, err := config.GetConfig(configFilePath)
	if err != nil {
		log.Fatalf("The config file is broken: %v", err.Error())
	}

	logger.SetupLogger(cfg.LogFile, cfg.LogLevel)
	log.Info("application started.")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()

	r, err := reminder.New(cfg.ConnectionString, cfg.RabbitMQ, ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Do()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("application is shutdown")
}
