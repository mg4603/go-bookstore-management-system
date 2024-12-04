package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mg4603/go-bookstore-management-system/pkg/controllers"
	"github.com/mg4603/go-bookstore-management-system/pkg/utils"
)

func RegisterBookstoreRoutes(r *mux.Router, controllers *controllers.BookstoreController) {
	r.Handle("/books/", utils.SetJSONContentType(http.HandlerFunc(controllers.CreateBook))).Methods("POST")
	r.Handle("/books/", utils.SetJSONContentType(http.HandlerFunc(controllers.GetBooks))).Methods("GET")
	r.Handle("/books/{id}", utils.SetJSONContentType(http.HandlerFunc(controllers.GetBookById))).Methods("GET")
	r.Handle("/books/{id}", utils.SetJSONContentType(http.HandlerFunc(controllers.DeleteBook))).Methods("DELETE")
	r.Handle("/books/{id}", utils.SetJSONContentType(http.HandlerFunc(controllers.UpdateBook))).Methods("PUT")
}
