package database

import (
    "gotest/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "log"
    // "os"
)

var DB *gorm.DB

func Connect() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file, using environment variables")
    }

    // ds= "host=" + os.Getenv("DB_HOST") +
    //     " user=" + os.Getenv("DB_USER") +
    //     " password=" + os.Getenv("DB_PASSWORD") +
    //     " dbname=" + os.Getenv("DB_NAME") +
    //     " port=" + os.Getenv("DB_PORT") +
    //     " sslmode=disable"n := "host=" + os.Getenv("DB_HOST") +
    //     " user=" + os.Getenv("DB_USER") +
    //     " password=" + os.Getenv("DB_PASSWORD") +
    //     " dbname=" + os.Getenv("DB_NAME") +
    //     " port=" + os.Getenv("DB_PORT") +
    //     " sslmode=disable"

    dsn := "host=" + "postgres" +
        " user=" + "postgres" +
        " password=" + "yourpassword" +
        " dbname=" + "yourdatabase" +
        " port=" + "5432" +
        " sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB = db

    // Auto-migrate the User model
    DB.AutoMigrate(&models.User{})
}
