package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		vars := mux.Vars(r)
		bookId, ok := vars["id"]

		if !ok {
			utils.HandleError(w, http.StatusBadRequest, "required field (id) is missing")
			return
		}

		ID, err := strconv.ParseInt(bookId, 0, 0)
		if err != nil {
			utils.HandleError(w, http.StatusBadRequest, "bad input: couldn't parse int id, from string bookId: "+bookId)
			return
		}

		bookDetails, err := db.GetBookById(ID)

		if err != nil {
			if err.Error() == fmt.Sprintf("book with ID %d not found", ID) {
				utils.HandleError(w, http.StatusNotFound, fmt.Sprintf("book with id %d does not exist in database", ID))
			} else {
				utils.HandleError(w, http.StatusInternalServerError, fmt.Sprintf("error occured while trying to fetch record from db: %s", err.Error()))
			}
			return
		}

		if err := json.NewEncoder(w).Encode(bookDetails); err != nil {
			utils.HandleError(w, http.StatusInternalServerError, fmt.Sprintf("error occurred while encoding response: %s", err.Error()))
			return
		}
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
