package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mg4603/go-bookstore-management-system/pkg/models"
	"github.com/mg4603/go-bookstore-management-system/pkg/utils"
)

type BookstoreController struct {
	CreateBook  http.HandlerFunc
	GetBooks    http.HandlerFunc
	GetBookById http.HandlerFunc
	UpdateBook  http.HandlerFunc
	DeleteBook  http.HandlerFunc
}

func NewBookStoreController(db *models.DBModel) *BookstoreController {
	return &BookstoreController{
		CreateBook:  CreateBookHandler(db),
		GetBooks:    GetBooksHandler(db),
		GetBookById: GetBookByIdHandler(db),
		UpdateBook:  UpdateBookHandler(db),
		DeleteBook:  DeleteBookHandler(db),
	}
}

func CreateBookHandler(db *models.DBModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetBooksHandler(db *models.DBModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newBooks, err := db.GetAllBooks()
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

func GetBookByIdHandler(db *models.DBModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func UpdateBookHandler(db *models.DBModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteBookHandler(db *models.DBModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
