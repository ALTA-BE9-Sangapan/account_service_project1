package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connect")
	}
	return db
}
