package controllers

import (
	"github.com/mg4603/go-bookstore-management-system/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setup() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Book{}); err != nil {
		return nil, err
	}
	return db, nil
}
