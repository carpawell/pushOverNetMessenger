package config

import (
	"errors"
	"flag"
)

var AppIdOrUserIdNotFound = errors.New("app and user flags are required")

type Config struct {
	Port   *string
	AppId  *string
	UserId *string
	DbHost *string
}

// Read command line arguments
func ReadConfig() (*Config, error) {
	var cfg = &Config{}
	cfg.Port = flag.String("port", "8080", "a port for service")
	cfg.AppId = flag.String("app", "", "application identifier")
	cfg.UserId = flag.String("user", "", "user device identifier")
	cfg.DbHost = flag.String("host", "", "database host")

	flag.Parse()

	if *cfg.AppId == "" || *cfg.UserId == "" {
		return nil, AppIdOrUserIdNotFound
	}

	return cfg, nil
}
