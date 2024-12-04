package config

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mockDB *gorm.DB

func mockOpen(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	mockDB := &gorm.DB{}
	return mockDB, nil
}
