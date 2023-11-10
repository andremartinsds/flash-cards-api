package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializeDatabase()

	if err != nil {
		errors.New("Database with error")
	}

	return nil
}

func GetDataSource() *gorm.DB {
	return db
}
