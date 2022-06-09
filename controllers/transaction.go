package controllers

import (
	"fmt"
	_entities "project1/entities"

	"database/sql"
)

func GetBalancebyPhone(db *sql.DB, phone string) ([]_entities.User, error) {
	fmt.Println("success")
	query, err := db.Query(`SELECT balance FROM user WHERE phone = ?`, phone)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	var user []_entities.User

	for query.Next() {
		var data _entities.User
		err := query.Scan(&data.Balance)

		if err != nil {
			fmt.Println("error2", err.Error())
		}
		user = append(user, data)
	}

	return user, nil
}
