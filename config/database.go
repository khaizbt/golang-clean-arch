package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

// SetupDB - setup dabase for sharing to all api
func init() {

	dsn := "root@tcp(127.0.0.1:3306)/go_shop?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err != nil {
		panic("migration failed")
	}

	db = database
}
