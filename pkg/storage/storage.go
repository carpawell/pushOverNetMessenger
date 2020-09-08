package storage

import (
	"github.com/carpawell/pushOverNetMessenger/pkg/config"
	"github.com/carpawell/pushOverNetMessenger/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Storage struct {
	Db *gorm.DB
}

// Creating connection to DataBase && auto migration
func New(cfg *config.Config) (*Storage, error) {
	dsn := utils.GetDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Notification{})
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, nil
}

// Add notification to DB
func (stg *Storage) AddNotification(msg string, sts int) {
	ntf := Notification{Message: msg, Status: sts, Time: time.Now()}
	stg.Db.Create(ntf)
}

// Get statistics for messages that are older then `from` argument
func (stg *Storage) GetMessageStatistics(from *time.Time) *MessagesStatistic {
	var succeed int64
	var failed int64
	stg.Db.Model(&Notification{}).Where("Time > ? AND status = ?", from, 1).Count(&succeed)
	stg.Db.Model(&Notification{}).Where("Time > ? AND status <> ?", from, 1).Count(&failed)
	return &MessagesStatistic{from, succeed + failed, succeed, failed}
}
