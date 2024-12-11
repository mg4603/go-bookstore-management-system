package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mg4603/go-bookstore-management-system/pkg/config"
	"github.com/mg4603/go-bookstore-management-system/pkg/controllers"
	"github.com/mg4603/go-bookstore-management-system/pkg/models"
	"github.com/mg4603/go-bookstore-management-system/pkg/routes"
	"gorm.io/gorm"
)

var db *models.DBModel

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func openDB(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	if db, err := gorm.Open(dialector, config); err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	} else {
		return db, nil
	}
}

func init() {
	config.Connect(openDB, loadEnv)
	bookDB := config.GetDB()
	if err := bookDB.AutoMigrate(&models.Book{}); err != nil {
		log.Printf("error during automigration: %s", err.Error())
		return
	}

	db = &models.DBModel{DB: bookDB}
}

func main() {
	bookstoreController := controllers.NewBookStoreController(db)

	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r, bookstoreController)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
