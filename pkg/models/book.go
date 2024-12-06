package models

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mg4603/go-bookstore-management-system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func openDB(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	if db, err := gorm.Open(dialector, config); err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	} else {
		return db, nil
	}
}

func init() {
	config.Connect(openDB, loadEnv)
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}
