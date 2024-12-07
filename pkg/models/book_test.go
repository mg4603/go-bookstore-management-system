package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setup() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&Book{}); err != nil {
		return nil, err
	}
	return db, nil
}

func TestCreateBook(t *testing.T) {

	mockDB, err := setup()
	assert.NoError(t, err, "failed to setup database: %w", err)

	defer func() {
		sqlDB, _ := mockDB.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}()

	tests := []struct {
		name          string
		book          *Book
		expectedError string
	}{
		{
			name: "Valid Book",
			book: &Book{
				Name:        "Book 1",
				Author:      "Author 1",
				Publication: "Publication 1",
			},
			expectedError: "",
		},
		{
			name: "Book missing field",
			book: &Book{
				Name:   "Book 2",
				Author: "Author 2",
			},
			expectedError: "missing required fields",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := CreateBook(tc.book, mockDB)
			if tc.expectedError != "" {
				assert.Error(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
