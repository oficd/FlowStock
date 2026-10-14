package database

import (
	"fmt"
	"log"
	"backend/config"
	"backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	// Load konfigurasi database dari .env
	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPass := config.GetEnv("DB_PASS", "12345678")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "5432")
	dbName := config.GetEnv("DB_NAME", "db_golang")

	// Format DSN untuk MySQL
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        dbHost,
        dbUser,
        dbPass,
        dbName,
        dbPort,
    )

	// Koneksi ke database
	var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    fmt.Println("Database connected successfully!")

    err = DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.Supplier{}, &models.User{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    fmt.Println("Database migrated successfully!")
}
