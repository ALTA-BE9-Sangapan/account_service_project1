package controllers

import (
	"errors"
	"fmt"
	"project1/entities"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser entities.User) error {
	result := db.Create(&newUser)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	return nil
}

func GetUserbyPassword(db *gorm.DB, password string) []entities.User {
	var users []entities.User

	tx := db.Find(&users, "password = ?", password)

	if tx.Error != nil {
		fmt.Println("error ", tx.Error)
	}

	return users
}

func GetUserbyID(db *gorm.DB, ID uint) []entities.User {
	var users []entities.User

	tx := db.Find(&users, "id = ?", ID)

	if tx.Error != nil {
		fmt.Println("error ", tx.Error)
	}

	return users
}
