package configs

import (
	h "go-jwt/app/helpers"
	m "go-jwt/app/models"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dbHost := h.GetEnv("DB_HOST")
	dbUser := h.GetEnv("DB_USER")
	dbPass := h.GetEnv("DB_PASS")
	dbName := h.GetEnv("DB_NAME")
	dbPort := h.GetEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&m.User{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}
