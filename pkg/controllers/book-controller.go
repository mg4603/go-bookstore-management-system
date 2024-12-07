package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type BookstoreController struct {
	DB *gorm.DB
}

func NewBookStoreController(db *gorm.DB) *BookstoreController {
	return &BookstoreController{DB: db}
}

func (c *BookstoreController) CreateBook(w http.ResponseWriter, r *http.Request) {

}

func (c *BookstoreController) GetBooks(w http.ResponseWriter, r *http.Request) {

}

func (c *BookstoreController) GetBookById(w http.ResponseWriter, r *http.Request) {

}

func (c *BookstoreController) UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func (c *BookstoreController) DeleteBook(w http.ResponseWriter, r *http.Request) {

}
