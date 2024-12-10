package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mg4603/go-bookstore-management-system/pkg/models"
	"github.com/mg4603/go-bookstore-management-system/pkg/tests"
	"github.com/mg4603/go-bookstore-management-system/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetBooksHandler(t *testing.T) {

	testTable := []struct {
		name           string
		mockSetup      func(db *models.DBModel)
		expectedStatus int
		expectedBody   string
	}{
		{name: "Successful retrieval of books",
			mockSetup: func(db *models.DBModel) {
				books := []models.Book{
					{Name: "Book1", Author: "Author1", Publication: "Publication1"},
					{Name: "Book2", Author: "Author2", Publication: "Publication2"},
				}

				for _, book := range books {
					err := db.CreateBook(&book)
					assert.NoError(t, err)
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"ID":1,"name":"Book1","author":"Author1","publication":"Publication1"},{"ID":2,"name":"Book2","author":"Author2","publication":"Publication2"}]`,
		},
		{name: "No book in db",
			mockSetup:      func(db *models.DBModel) {},
			expectedStatus: http.StatusOK,
			expectedBody:   `[]`,
		},
		{name: "Database error",
			mockSetup: func(db *models.DBModel) {
				sqlDB, _ := db.DB.DB()
				if sqlDB != nil {
					sqlDB.Close()
				}
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, err := tests.Setup()
			assert.NoError(t, err, "unexpected error while setting up mock database: %w", err)
			defer func() {
				sqlDB, _ := mockDB.DB()
				if sqlDB != nil {
					sqlDB.Close()
				}
			}()
			db := &models.DBModel{DB: mockDB}
			tt.mockSetup(db)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/books/", nil)

			handler := utils.SetJSONContentType(GetBooksHandler(db))
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestGetBookByIdHandler(t *testing.T) {
	testCases := []struct {
		name           string
		bookId         string
		mockSetup      func(db *models.DBModel)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "Successful retrieval of book",
			bookId: "1",
			mockSetup: func(db *models.DBModel) {
				book := models.Book{Name: "Book1", Author: "Author1", Publication: "Publication1"}
				err := db.CreateBook(&book)
				assert.NoError(t, err)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"ID":1,"name":"Book1","author":"Author1","publication":"Publication1"}`,
		},
		{
			name:   "Book Not found",
			bookId: "9999",
			mockSetup: func(db *models.DBModel) {

			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
		{
			name:   "Invalid book ID",
			bookId: "abc",
			mockSetup: func(db *models.DBModel) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, err := tests.Setup()
			assert.NoError(t, err)
			defer func() {
				sqlDB, _ := mockDB.DB()
				if sqlDB != nil {
					sqlDB.Close()
				}
			}()
			db := &models.DBModel{DB: mockDB}
			tt.mockSetup(db)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/books/{id}", nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.bookId})

			handler := utils.SetJSONContentType(GetBookByIdHandler(db))
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestCreateBookHandler(t *testing.T) {
	testCases := []struct {
		name           string
		inputBody      interface{}
		mockSetup      func(db *models.DBModel)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Successful book creation",
			inputBody: &models.Book{Name: "Book1", Author: "Author1", Publication: "Publication1"},
			mockSetup: func(db *models.DBModel) {

			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"ID":1,"name":"Book1","author":"Author1","publication":"Publication1"}`,
		},
		{
			name:      "Invalid request body",
			inputBody: map[string]string{"invalid field": "invalid value"},
			mockSetup: func(db *models.DBModel) {

			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
		{
			name:      "Database error",
			inputBody: models.Book{Name: "Book2", Author: "Author2", Publication: "Publication2"},
			mockSetup: func(db *models.DBModel) {
				sqlDB, _ := db.DB.DB()
				if sqlDB != nil {
					sqlDB.Close()
				}
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, err := tests.Setup()
			assert.NoError(t, err)
			defer func() {
				sqlDB, _ := mockDB.DB()
				if sqlDB != nil {
					sqlDB.Close()
				}
			}()
			db := &models.DBModel{DB: mockDB}
			tt.mockSetup(db)

			bodyBytes, err := json.Marshal(tt.inputBody)
			assert.NoError(t, err, "failed to marshal input body")

			req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(bodyBytes))
			rec := httptest.NewRecorder()

			handler := utils.SetJSONContentType(CreateBookHandler(db))
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestDeleteBookHandler(t *testing.T) {
	testCases := []struct {
		name           string
		bookId         string
		mockSetup      func(db *models.DBModel)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "Successful deletion",
			bookId: "1",
			mockSetup: func(db *models.DBModel) {
				book := &models.Book{Name: "Book1", Author: "Author1", Publication: "Publication1"}
				err := db.CreateBook(book)
				assert.NoError(t, err)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"ID":1,"name":"Book1","author":"Author1","publication":"Publication1"}`,
		},
		{
			name:           "Missing id parameter",
			bookId:         "",
			mockSetup:      func(db *models.DBModel) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
		{
			name:           "Invalid id format",
			bookId:         "abc",
			mockSetup:      func(db *models.DBModel) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
		{
			name:           "Book not found",
			bookId:         "9999",
			mockSetup:      func(db *models.DBModel) {},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
		{
			name:   "Database error",
			bookId: "1",
			mockSetup: func(db *models.DBModel) {
				sqlDB, _ := db.DB.DB()
				if sqlDB != nil {
					sqlDB.Close()
				}
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"An error occurred. Please try again later."}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDB, err := tests.Setup()
			assert.NoError(t, err)

			db := &models.DBModel{DB: mockDB}

			tc.mockSetup(db)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/books/{id}", nil)
			req = mux.SetURLVars(req, map[string]string{"id": tc.bookId})

			handler := utils.SetJSONContentType(DeleteBookHandler(db))
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedBody, rec.Body.String())
		})
	}
}
