package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() error {
	if err := godotenv.Load(); err != nil {
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
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	fmt.Println("Database connection established!")
	return nil
}
