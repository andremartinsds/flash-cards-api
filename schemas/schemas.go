package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name  string
	email string
	pass  string
	Dec   []Dec
}

type Dec struct {
	gorm.Model
	name       string
	FlashCards []FlashCards
	UserID     uint
}

type FlashCards struct {
	gorm.Model
	front        string
	back         string
	nextRevision time.Time
	lastRevision time.Time
	FlashCardsID uint
}
