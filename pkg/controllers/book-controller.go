package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mg4603/go-bookstore-management-system/pkg/models"
	"github.com/mg4603/go-bookstore-management-system/pkg/utils"
	"gorm.io/gorm"
)

type BookstoreController struct {
	CreateBook  http.HandlerFunc
	GetBooks    http.HandlerFunc
	GetBookById http.HandlerFunc
	UpdateBook  http.HandlerFunc
	DeleteBook  http.HandlerFunc
}

func NewBookStoreController(db *gorm.DB) *BookstoreController {
	return &BookstoreController{
		CreateBook:  CreateBookHandler(db),
		GetBooks:    GetBooksHandler(db),
		GetBookById: GetBookByIdHandler(db),
		UpdateBook:  UpdateBookHandler(db),
		DeleteBook:  DeleteBookHandler(db),
	}
}

func CreateBookHandler(db *gorm.DB) http.HandlerFunc {

}

func GetBooksHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newBooks, err := models.GetAllBooks(db)
		if err != nil {
			utils.HandleError(w, http.StatusInternalServerError, "error fetching books from database")
			return
		}

		if err := json.NewEncoder(w).Encode(newBooks); err != nil {
			utils.HandleError(w, http.StatusInternalServerError, "error marshalling new books")
			return
		}
	}
}

func GetBookByIdHandler(db *gorm.DB) http.HandlerFunc {

}

func UpdateBookHandler(db *gorm.DB) http.HandlerFunc {

}

func DeleteBookHandler(db *gorm.DB) http.HandlerFunc {

}
