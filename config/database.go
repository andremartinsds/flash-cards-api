package config

import (
	"errors"
	"os"

	"github.com/andremartinsds/flash-cards-api/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_ROOT_PASSWORD") + "@tcp(" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + ")/" + os.Getenv("MYSQL_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
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
