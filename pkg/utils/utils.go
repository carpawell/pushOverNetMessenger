package utils

import (
	"github.com/carpawell/pushOverNetMessenger/pkg/constants"
	"os"
)

func GetDSN() string {
	envDSN := os.Getenv("DB_DSN")

	if envDSN == "" {
		return constants.DefaultDbURI
	}

	return envDSN
}
