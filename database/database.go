package database

import (
	"Assignment_2/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Debug
	db = db.Debug()

	// Open Database
	autoMigrate()
}

func GetDb() *gorm.DB {
	return db
}

func autoMigrate() {
	db.AutoMigrate(&models.Order{}, &models.Item{})
}
