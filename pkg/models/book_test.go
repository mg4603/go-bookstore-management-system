package models

import (
	"errors"
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

func TestGetAllBooks(t *testing.T) {
	mockDB, err := setup()
	assert.NoError(t, err, "failed to setup database: %w", err)

	defer func() {
		sqlDB, _ := mockDB.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}()
	tests := []struct {
		name           string
		setupBooks     []Book
		expectedBooks  []Book
		expectedLength int
	}{
		{
			name:           "No book in database",
			setupBooks:     nil,
			expectedBooks:  []Book{},
			expectedLength: 0,
		},
		{
			name: "Single book in database",
			setupBooks: []Book{
				{Name: "Book 1", Author: "Author 1", Publication: "Publication 1"},
			},
			expectedBooks: []Book{
				{Name: "Book 1", Author: "Author 1", Publication: "Publication 1"},
			},
			expectedLength: 1,
		},
		{
			name: "Multiple books in database",
			setupBooks: []Book{
				{Name: "Book 1", Author: "Author 1", Publication: "Publication 1"},
				{Name: "Book 2", Author: "Author 1", Publication: "Publication 1"},
			},
			expectedBooks: []Book{
				{Name: "Book 1", Author: "Author 1", Publication: "Publication 1"},
				{Name: "Book 2", Author: "Author 1", Publication: "Publication 1"},
			},
			expectedLength: 2,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB.Exec("DELETE FROM books")

			for _, book := range tc.setupBooks {
				err := CreateBook(&book, mockDB)
				assert.NoError(t, err, "failed to insert setup data: %w", err)
			}

			books, err := GetAllBooks(mockDB)
			assert.NoError(t, err, "error retrieving books: %w", err)
			assert.Equal(t, len(books), tc.expectedLength, "incorrect nmber of books returned")

			for i, book := range tc.expectedBooks {
				assert.Equal(t, book.Name, books[i].Name, "book name wanted = %v; got %v", book.Name, books[i].Name)
				assert.Equal(t, book.Author, books[i].Author, "book author wanted = %v; got %v", book.Author, books[i].Author)
				assert.Equal(t, book.Publication, books[i].Publication, "book publication wanted = %v; got %v", book.Publication, books[i].Publication)
			}
		})
	}
}

func TestGetBookById(t *testing.T) {
	mockDB, err := setup()
	assert.NoError(t, err, "failed to setup test database")

	defer func() {
		sqlDB, _ := mockDB.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}()

	seedBooks := []Book{
		{Name: "Name 1", Author: "Author 1", Publication: "Publication 1"},
		{Name: "Name 2", Author: "Author 2", Publication: "Publication 2"},
	}

	for _, book := range seedBooks {
		err := CreateBook(&book, mockDB)
		assert.NoError(t, err, "failed to seed database")
	}

	tests := []struct {
		name          string
		bookID        int64
		expectedBook  *Book
		expectedError error
	}{
		{
			name:          "Valid book ID 1",
			bookID:        1,
			expectedBook:  &Book{Name: "Name 1", Author: "Author 1", Publication: "Publication 1"},
			expectedError: nil,
		},
		{
			name:          "Valid book ID 2",
			bookID:        2,
			expectedBook:  &Book{Name: "Name 2", Author: "Author 2", Publication: "Publication 2"},
			expectedError: nil,
		},
		{
			name:          "Non-existent book ID",
			bookID:        9999999,
			expectedBook:  nil,
			expectedError: errors.New("book with ID 9999999 not found"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			book, err := GetBookById(tc.bookID, mockDB)

			if tc.expectedError != nil {
				assert.Error(t, err, "expected error got none")
				assert.EqualError(t, err, tc.expectedError.Error(), "expected error = %w; got %w", tc.expectedError, err)
			} else {
				assert.NoError(t, err, "unexpected error: %w", err)
			}

			if tc.expectedBook != nil {
				assert.NotNil(t, book, "expected book but got nil")
				assert.Equal(t, tc.expectedBook.Name, book.Name, "expected book name: %v; got %v", tc.expectedBook.Name, book.Name)
				assert.Equal(t, tc.expectedBook.Author, book.Author, "expected author name: %v; got %v", tc.expectedBook.Author, book.Author)
				assert.Equal(t, tc.expectedBook.Publication, book.Publication, "expected publication name: %v; got %v", tc.expectedBook.Publication, book.Publication)
			} else {
				assert.Nil(t, book, "expected nil book but got one %v", book)
			}
		})
	}
}
