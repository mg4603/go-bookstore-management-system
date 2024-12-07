package models

import (
	"errors"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Author      string `gorm:"not null" json:"author"`
	Publication string `gorm:"not null" json:"publication"`
}

func CreateBook(b *Book, db *gorm.DB) error {
	if b.Author == "" || b.Name == "" || b.Publication == "" {
		return errors.New("missing required fields")
	}
	if result := db.Create(b); result.Error != nil {
		return result.Error
	}
	return nil
}
