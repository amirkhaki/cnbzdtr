package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Token string `env:"CD_DISCORD_TOKEN"`
}

func parseEnv(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("Could not parse config from env: %w", err)
	}
	return nil
}

func New() (Config, error) {
	cfg := Config{}
	err := parseEnv(&cfg)
	return cfg, err

}
