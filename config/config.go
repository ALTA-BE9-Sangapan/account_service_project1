package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbConnection := os.Getenv("DB_CONNECTION")

	db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
