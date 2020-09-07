package storage

import (
	"github.com/carpawell/pushOverNetMessenger/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
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

func (stg *Storage) AddNotification(msg string, sts int) {
	ntf := Notification{Message: msg, Status: sts, Time: time.Now()}
	stg.Db.Create(ntf)
}

func (stg *Storage) GetMessageStatistics(from *time.Time) *MessagesStatistic {
	var succeed int64
	var failed int64
	stg.Db.Model(&Notification{}).Where("Time > ? AND status = ?", from.Format("2006-01-02 15:04:05"), 1).Count(&succeed)
	stg.Db.Model(&Notification{}).Where("Time > ? AND status <> ?", from.Format("2006-01-02 15:04:05"), 1).Count(&failed)
	return &MessagesStatistic{from, succeed + failed, succeed, failed}
}
