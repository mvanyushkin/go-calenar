package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/mvanyushkin/go-calendar/calendar"
	"github.com/mvanyushkin/go-calendar/calendar/store"
	"github.com/mvanyushkin/go-calendar/config"
	server "github.com/mvanyushkin/go-calendar/grpc"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		fmt.Printf("The config file is broken: %v", err.Error())
		os.Exit(-1)
	}

	setupLogger(cfg.LogFile, cfg.LogLevel)
	log.Info("application started.")
	serve(err, cfg.HttpListen)
}

func serve(err error, binding string) {
	lis, err := net.Listen("tcp", binding)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server.RegisterCalendarServer(grpcServer, server.CalendarHandler{
		Calendar: calendar.NewCalendar(store.NewInMemoryEventStore()),
	})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Errorf("the serve process has failed, occurred an exception: %v \n", err.Error())
	}
}

func getConfig() (*config.Config, error) {
	configFilePath := flag.String("config", "", "settings file")
	flag.Parse()

	loader := confita.NewLoader(
		env.NewBackend(),
		file.NewBackend(*configFilePath),
	)

	cfg := config.Config{}
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setupLogger(filePath string, logLevel string) {
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
