package models

import (
	"readinglist-api/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model        // using gorm convention for auto id, created_at, updated_at, deleted_at
	Title      string `json:"title"`
	Author     string `json:"author"`
	Category   string `json:"category"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
