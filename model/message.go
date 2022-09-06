package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Text   string
	UserID int
	User   User
}
