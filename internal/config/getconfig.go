package config

import (
	"context"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"os"
	"path"
)

func GetConfig(configFilePath *string) (*Config, error) {
	_, err := os.Stat(*configFilePath)

	if configFilePath == nil || os.IsNotExist(err) {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		defaultConfigPath := path.Join(wd, "local_config.json")
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
