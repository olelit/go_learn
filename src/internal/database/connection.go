package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"learn/go/internal/models"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=db user=go_local password=go_local dbname=go_local port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB from gorm.DB:", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("database ping failed:", err)
	}

	err = db.AutoMigrate(&models.Users{})
	if err != nil {
		log.Fatal("migration failed:", err)
	}

	DB = db
	log.Println("Database connected and migrated successfully")
}
