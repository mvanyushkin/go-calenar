package main

import (
	"context"
	"github.com/mvanyushkin/go-calendar/internal/config"
	"github.com/mvanyushkin/go-calendar/internal/sender"
	"github.com/mvanyushkin/go-calendar/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.GetConfig()
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

	s := sender.CreateSender(cfg.RabbitMQ, ctx)
	err = s.ListenMessages()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("App is shutdown")
	}
}
