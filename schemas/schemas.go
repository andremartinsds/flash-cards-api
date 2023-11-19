package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Pass  string
	Dec   []Dec
}

func (u *User) UserExists() bool {
	return u.ID != 0
}

type Dec struct {
	gorm.Model
	Name       string
	FlashCards []FlashCards
	UserID     uint
}

func (d *Dec) DecExists() bool {
	return d.ID != 0
}

type FlashCards struct {
	gorm.Model
	Front        string
	Back         string
	NextRevision time.Time
	LastRevision time.Time
	DecID        uint
}

func (f *FlashCards) FlashCardExists() bool {
	return f.ID != 0
}
