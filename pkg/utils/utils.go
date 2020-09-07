package utils

import (
	"github.com/carpawell/pushOverNetMessenger/pkg/config"
	"github.com/carpawell/pushOverNetMessenger/pkg/constants"
	"os"
	"strings"
	"time"
)

func GetDSN(config *config.Config) string {
	if config != nil && *config.DbHost != "" {
		return strings.Replace(constants.DefaultDbURI, "localhost", *config.DbHost, 1)
	}

	envDSN := os.Getenv("DB_DSN")

	if envDSN == "" {
		return constants.DefaultDbURI
	}

	return envDSN
}

func ParseTime(str string) (*time.Time, error) {
	parsedTime, err := time.Parse(constants.TimeLayout, str)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}
