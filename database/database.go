package database

import (
	"Assignment_2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// To Do

	// DSN
	db, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/dts_golang?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
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
