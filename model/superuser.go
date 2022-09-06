package model

import (
	"gorm.io/gorm"
)

type SuperUser struct {
	gorm.Model
	UserID int
	User   User
}
