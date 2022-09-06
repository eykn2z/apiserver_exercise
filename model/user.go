package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	GroupID int //must for relation
	Group   Group
}
