package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dbConnection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connect")
	}
	return db
}
