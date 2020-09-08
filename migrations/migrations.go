package main

import (
	"github.com/carpawell/pushOverNetMessenger/pkg/storage"
	"github.com/carpawell/pushOverNetMessenger/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Migration if needed. Not Part of application. Reading DB_DSN env.
func main() {
	dsn := utils.GetDSN(nil)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database:%s\n", err)
	}

	err = db.AutoMigrate(&storage.Notification{})
	if err != nil {
		log.Fatalf("migrations:%s\n", err)
	}
}
