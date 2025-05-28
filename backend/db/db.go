package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "backend/models"
)

var DB *gorm.DB

func Init() {
    dsn := "host=localhost user=postgres password=SandorIST8642 dbname=hotel_db port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }
    DB.AutoMigrate(&models.Hotel{}, &models.Reservation{})
}
