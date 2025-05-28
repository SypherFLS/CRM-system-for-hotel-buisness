package db

import (
    "log"
    "backend/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    dsn := "host=localhost user=postgres password=SandorIST8642 dbname=hotel_db port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    err = DB.AutoMigrate(
        &models.Room{},
        &models.Client{},
        &models.Employee{},
        &models.CleaningSchedule{},
    )
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
}
