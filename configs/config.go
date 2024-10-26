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
	Grpc struct {
		Host           string `yaml:"HOST" env:"GRPC_HOST" env-description:"GRPC HOST" env-default:"localhost"`
		Port           string `yaml:"PORT" env:"GRPC_PORT" env-description:"GRPC PORT" env-default:"50051"`
		Timeout        int    `yaml:"TIMEOUT" env:"GRPC_CONTEXT_TIMEOUT" env-description:"GRPC CONTEXT TIMEOUT" env-default:"30000"`
		MaxGrpcReceive int    `yaml:"RECEIVE" env:"MAX_GRPC_RECEIVE" env-description:"MAX GRPC RECEIVE" env-default:"1073741824"`
		MaxGrpcSend    int    `yaml:"RECEIVE" env:"MAX_GRPC_RECEIVE" env-description:"MAX GRPC RECEIVE" env-default:"1073741824"`
	} `yaml:"GRPC"`
	LoggerConfig struct {
		Environment string `yaml:"ENVIRONMENT" env:"LOGGER_ENVIRONMENT"`
		Level       string `yaml:"LEVEL" env:"LOGGER_LEVEL"`
	} `yaml:"LOGGER"`
}

func NewConfig() (*Config, error) {
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
