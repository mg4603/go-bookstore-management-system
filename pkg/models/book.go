package models

import (
	"errors"
	"fmt"

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

func GetAllBooks(db *gorm.DB) ([]Book, error) {
	var books []Book
	if result := db.Find(&books); result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func GetBookById(id int64, db *gorm.DB) (*Book, error) {
	var book Book

	if result := db.First(&book, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("book with ID %d not found", id)
		}
		return nil, result.Error
	}
	return &book, nil
}
