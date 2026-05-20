package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"restapis-go/internal/entities"
)

func Connect() {
	// Connect with gorm
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, ".ENV not found, error: %v\n", err)
		os.Exit(1)
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB object: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Successfully connect to Neon Postgres Database!")

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database migrated successfully!")
}