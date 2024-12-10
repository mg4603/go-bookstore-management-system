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
		createBook := &models.Book{}

		if err := utils.ParseBody(r, createBook); err != nil {
			utils.HandleError(w, http.StatusBadRequest, fmt.Sprintf("error parsing input into book model: %s", err))
			return
		}

		if err := db.CreateBook(createBook); err != nil {
			utils.HandleError(w, http.StatusInternalServerError, fmt.Sprintf("error while trying to create book: %s", err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(createBook); err != nil {
			utils.HandleError(w, http.StatusInternalServerError, fmt.Sprintf("error occured while encoding created book: %s", err.Error()))
			return
		}
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
		vars := mux.Vars(r)
		bookId, ok := vars["id"]
		if !ok {
			utils.HandleError(w, http.StatusBadRequest, "required field (id) is missing")
			return
		}

		ID, err := strconv.ParseInt(bookId, 0, 0)

		if err != nil {
			utils.HandleError(w, http.StatusBadRequest, fmt.Sprintf("bad input; couldn't parse integer id from string bookID: %s", err.Error()))
			return
		}

		book, err := db.DeleteBook(ID)
		if err != nil {
			if err.Error() == fmt.Sprintf("book with ID %d not found", ID) {
				utils.HandleError(w, http.StatusNotFound, fmt.Sprintf("book with id %d does not exist in database", ID))
			} else {
				utils.HandleError(w, http.StatusInternalServerError, fmt.Sprintf("err while trying to delete book of id %d from db: %s", ID, err.Error()))
			}
			return
		}

		if err := json.NewEncoder(w).Encode(&book); err != nil {
			utils.HandleError(w, http.StatusInternalServerError, fmt.Sprintf("error occurred while encoding server response: %s", err.Error()))
			return
		}

	}
}
