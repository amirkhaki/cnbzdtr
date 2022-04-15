package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Token        string `env:"CD_DISCORD_TOKEN"`
	Redis_DSN    string `env:"CD_REDIS_DSN"`
	Cipher_key   string `env:"CD_CIPHER_KEY"`
	Port         string `env:"CD_HTTP_PORT"`
	MessageScore uint64 `env:"CD_MESSAGE_SCORE"`
	InviteScore  uint64 `env:"CD_INVITE_SCORE"`
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
	if err == nil {
		if len(cfg.Cipher_key)%16 != 0 {
			err = fmt.Errorf("cipher key length should be 16 or 32")
		}
	}
	return cfg, err

}
