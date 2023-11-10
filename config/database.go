package config

import (
	"errors"

	"github.com/andremartinsds/flash-cards-api/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	//TODO: get values from .env
	dsn := "flash_cards_api_user:flash_cards_api_user_pass@tcp(localhost:3313)/flash_cards_api_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New("Database with error")
	}

	err = db.AutoMigrate(&schemas.User{}, &schemas.Dec{}, &schemas.FlashCards{})

	if err != nil {
		return nil, errors.New("Database error at automigrate")
	}

	return db, nil
}
