package config

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type GHPrivateKeyField []byte

func (f *GHPrivateKeyField) SetValue(s string) error {
	privateKey, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		return fmt.Errorf("gitHubPrivateKey should be base64 decoded")
	}
	*f = privateKey
	return nil
}

type Config struct {
	PGUsername   string            `env:"PG_USERNAME" env-required:"true"`
	PGPassword   string            `env:"PG_PASSWORD" env-required:"true"`
	PGDBName     string            `env:"PG_DB_NAME" env-required:"true"`
	GHPrivateKey GHPrivateKeyField `env:"GH_PRIVATE_KEY" env-required:"true"`
	GHAppID      int64             `env:"GH_APP_ID" env-required:"true"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
