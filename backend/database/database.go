package database

import (
    "gotest/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "log"
    "os"
)

var DB *gorm.DB

func Connect() {
    if err := godotenv.Load(); err != nil {
        log.Println("⛔ ERROR", err)
    }

    dsn := "host=" + os.Getenv("DB_HOST") +
        " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " port=" + os.Getenv("DB_PORT") +
        " sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("⛔ ERROR", err)
    }
    DB = db

    // Auto-migrate the User model
    DB.AutoMigrate(&models.User{})
}
