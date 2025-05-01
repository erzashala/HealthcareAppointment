package database

import (
    "log"
    "healthcare-api/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("healthcare.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Could not connect to database")
    }

    DB.AutoMigrate(&models.Patient{}, &models.Appointment{})
}