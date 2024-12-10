package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BookstoreDB interface {
	CreateBook(b *Book) error
	GetAllBooks() ([]Book, error)
	GetBookById(id int64) (*Book, error)
	DeleteBook(id int64) (*Book, error)
}

type Book struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `gorm:"index" json:"-"`
	Name        string    `gorm:"not null" json:"name"`
	Author      string    `gorm:"not null" json:"author"`
	Publication string    `gorm:"not null" json:"publication"`
}

type DBModel struct {
	DB *gorm.DB
}

func (db *DBModel) CreateBook(b *Book) error {
	if b.Author == "" || b.Name == "" || b.Publication == "" {
		return errors.New("missing required fields")
	}
	if result := db.DB.Create(b); result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DBModel) GetAllBooks() ([]Book, error) {
	var books []Book
	if result := db.DB.Find(&books); result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (db *DBModel) GetBookById(id int64) (*Book, error) {
	var book Book

	if result := db.DB.First(&book, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("book with ID %d not found", id)
		}
		return nil, result.Error
	}
	return &book, nil
}

func (db *DBModel) DeleteBook(id int64) (*Book, error) {
	var book Book
	if result := db.DB.First(&book, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("book with ID %d not found", id)
		}
		return nil, result.Error
	}

	if result := db.DB.Delete(&book); result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}
