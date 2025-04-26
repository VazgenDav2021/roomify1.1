package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

// Инициализация базы данных
func ConnectDatabase() {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Подключаемся к базе данных
	DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully")
}
