package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func CreateBook(b *Book, db *gorm.DB) error {
	if result := db.Create(b); result.Error != nil {
		return result.Error
	}
	return nil
}
