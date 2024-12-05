package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type DBOpener func(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error)

type EnvLoader func() error

func Connect(opener DBOpener, loader EnvLoader) error {
	if err := loader(); err != nil {
		return fmt.Errorf("error while loading .env file: %w", err)
	}

	dbUserName := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbUserName == "" || dbPassword == "" || dbName == "" {
		return fmt.Errorf("missing one or more environment variables: DB_USER_NAME, DB_PASSWORD, DB_NAME")
	}
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		dbUserName, dbPassword, dbName)

	var err error
	db, err = opener(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	fmt.Println("Database connection established!")
	return nil
}

func GetDB() *gorm.DB {
	return db
}
