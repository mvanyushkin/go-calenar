package config

import (
	"context"
	"flag"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"os"
	"path"
)

func GetConfig() (*Config, error) {
	defaultConfigFileName := "local_config.json"
	configFilePath := flag.String("config", "", "settings file")
	flag.Parse()

	_, err := os.Stat(*configFilePath)

	if configFilePath == nil || os.IsNotExist(err) {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		defaultConfigPath := path.Join(wd, defaultConfigFileName)
		configFilePath = &defaultConfigPath
	}

	loader := confita.NewLoader(
		env.NewBackend(),
		file.NewBackend(*configFilePath),
	)

	cfg := Config{}
	err = loader.Load(context.Background(), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
