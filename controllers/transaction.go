package controllers

import (
	"fmt"
	_entities "project1/entities"

	"database/sql"
)

func GetBalancebyPhone(db *sql.DB, phone string) ([]_entities.User, error) {

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

func TopUp(db *sql.DB, topup int, phone string) error {
	_, err := db.Query(`INSERT INTO topup (TopUp_phone, TopUp_balance) VALUES (?, ?)`, phone, topup)

	if err != nil {
		fmt.Println("error1", err.Error())
		return err
	}
	// test, err1 := db.Query(`SELECT Balance FROM user WHERE Phone = ? UNION ALL SELECT TopUp_Balance FROM topup WHERE TopUp_phone = ?`, phone, phone)
	_, err1 := db.Query(`UPDATE user SET Balance = (SELECT Balance WHERE Phone = ?) + (SELECT TopUp_balance FROM topup ORDER BY Created_At DESC LIMIT 1) `, phone)
	if err1 != nil {
		fmt.Println("error2", err1.Error())
		return err1
	} else {
		return nil
	}
}

func GetNewBalance(db *sql.DB, phone string) ([]_entities.User, error) {
	query, err := db.Query(`SELECT balance from user WHERE phone = ? AND phone is not null`, phone)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	var other []_entities.User

	for query.Next() {
		var data _entities.User
		err := query.Scan(&data.Balance)

		if err != nil {
			fmt.Println("error2", err.Error())
		}
		other = append(other, data)
	}
	return other, nil
}
