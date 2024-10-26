package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/multierr"
)

type Config struct {
	Http struct {
		Host string `yaml:"HOST" env:"HTTP_HOST" env-description:"HTTP HOST" env-default:"localhost"`
		Port string `yaml:"PORT" env:"HTTP_PORT" env-description:"HTTP PORT" env-default:":8080"`
	} `yaml:"HTTP"`
	In struct {
		Count int `yaml:"COUNT" env:"IN_COUNT" env-description:"count of IN ports"`
	} `yaml:"IN"`
	Out struct {
		Count int `yaml:"COUNT" env:"OUT_COUNT" env-description:"count of OUT ports"`
	} `yaml:"OUT"`
}

func New() (*Config, error) {
	var (
		errorBuilder error
		cfg          Config
	)

	err := cleanenv.ReadConfig("./configs/config.yml", &cfg)
	multierr.AppendInto(&errorBuilder, err)

	if errorBuilder != nil {
		return nil, errorBuilder
	}

	return &cfg, nil
}
