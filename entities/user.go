package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Phone    string
	Password string `gorm:"unique"`
	Balance  int
	Gender   string
	Address  string
	ID       string
}
