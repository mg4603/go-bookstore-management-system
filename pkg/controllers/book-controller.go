package controllers

import (
	"net/http"

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

}

func GetBookByIdHandler(db *gorm.DB) http.HandlerFunc {

}

func UpdateBookHandler(db *gorm.DB) http.HandlerFunc {

}

func DeleteBookHandler(db *gorm.DB) http.HandlerFunc {

}
