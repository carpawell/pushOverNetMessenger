package storage

import (
	"github.com/carpawell/pushOverNetMessenger/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Db *gorm.DB
}

func New() (*Storage, error) {
	dsn := utils.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, nil
}
