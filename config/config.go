package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	err = loadEnv()

	if err != nil {
		return err
	}

	db, err = InitializeDatabase()

	if err != nil {
		return err
	}

	return nil
}

func GetDataSource() *gorm.DB {
	return db
}
