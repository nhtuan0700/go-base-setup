package configs

import (
	"fmt"

	env "github.com/caarlos0/env/v11"
)

type Config struct {
	Log      Log
	Database Database
	HTTP     HTTP
}

func NewConfig() (Config, error) {
	var (
		config = Config{}
		err    error
	)

	err = env.Parse(&config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse env: %w", err)
	}

	return config, nil
}
