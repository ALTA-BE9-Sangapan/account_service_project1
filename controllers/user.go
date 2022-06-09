package controllers

import (
	"fmt"
	_entities "project1/entities"

	"database/sql"
)

func CreateUser(db *sql.DB, newUser _entities.User) error {

	var query = (`INSERT INTO user (Name, Phone, Password, Balance, Gender, Address) VALUES (?, ?, ?, 0, ?, ?)`)
	insert, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		return errPrepare
	}

	_, err := insert.Exec(newUser.Name, newUser.Phone, newUser.Password, newUser.Gender, newUser.Address)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetUserbyPhone(db *sql.DB, phone string, password string) ([]_entities.User, error) {
	query, err := db.Query(`SELECT name, phone, balance, gender, address FROM user WHERE phone = ? AND password = ?`, phone, password)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	var user []_entities.User

	for query.Next() {
		var data _entities.User
		err := query.Scan(&data.Name, &data.Phone, &data.Balance, &data.Gender, &data.Address)

		if err != nil {
			fmt.Println("error2", err.Error())
		}
		user = append(user, data)
	}

	return user, nil
}

func UpdateName(db *sql.DB, name string, phone string) error {
	query, err := db.Prepare(`UPDATE user SET name = ? WHERE phone = ?`)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	_, err1 := query.Exec(name, phone)

	if err1 != nil {
		return err1
	} else {
		return nil
	}

}

func UpdateGender(db *sql.DB, gender string, phone string) error {
	query, err := db.Prepare(`UPDATE user SET gender = ? WHERE phone = ?`)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	_, err1 := query.Exec(gender, phone)

	if err1 != nil {
		return err1
	} else {
		return nil
	}

}

func UpdateAddress(db *sql.DB, address string, phone string) error {
	query, err := db.Prepare(`UPDATE user SET address = ? WHERE phone = ?`)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	_, err1 := query.Exec(address, phone)

	if err1 != nil {
		return err1
	} else {
		return nil
	}

}

func DeleteAccount(db *sql.DB, phone string) error {
	query, err := db.Prepare(`DELETE from user WHERE phone = ?`)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	_, err1 := query.Exec(phone)

	if err1 != nil {
		return err1
	} else {
		return nil
	}
}
