package handler

import (
	"github.com/andremartinsds/flash-cards-api/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitializeHandler() {
	DB = config.GetDataSource()
}
