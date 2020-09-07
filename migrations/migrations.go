package main

import (
	"github.com/carpawell/pushOverNetMessenger/pkg/storage"
	"github.com/carpawell/pushOverNetMessenger/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := utils.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database:%s\n", err)
	}

	err = db.AutoMigrate(&storage.Notification{})
	if err != nil {
		log.Fatalf("migrations:%s\n", err)
	}
}
