package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}
