package models

import (
	"testproject/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
