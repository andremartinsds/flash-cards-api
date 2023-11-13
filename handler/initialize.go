package handler

import (
	"github.com/andremartinsds/flash-cards-api/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetDataSource()
}
