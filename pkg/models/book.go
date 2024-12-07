package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Author      string `gorm:"not null" json:"author"`
	Publication string `gorm:"not null" json:"publication"`
}
