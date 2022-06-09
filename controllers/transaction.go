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
	_, err := db.Query(`INSERT INTO topup (TopUpPhone, TopUpBalance) VALUES (?, ?)`, phone, topup)

	if err != nil {
		fmt.Println("error1", err.Error())
		return err
	}

	_, err1 := db.Query(`UPDATE user SET Balance = (SELECT Balance WHERE Phone = ?) + (SELECT TopUpBalance FROM topup ORDER BY CreatedAt DESC LIMIT 1)`, phone)
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

func Transfer(db *sql.DB, receiver string, transfer int, phone string) error {
	_, err := db.Query(`INSERT INTO transfer (SenderPhone, ReceiverPhone, TransferBalance) VALUES (?, ?, ?)`, phone, receiver, transfer)

	if err != nil {
		fmt.Println("error1", err.Error())
		return err
	} else {
		_, err1 := db.Query(`UPDATE user SET Balance = (SELECT Balance WHERE Phone = ?) - (SELECT TransferBalance FROM transfer ORDER BY CreatedAt DESC LIMIT 1)`, phone)
		if err1 != nil {
			fmt.Println("error2", err1.Error())
			return err1
		} else {
			_, err2 := db.Query(`UPDATE user SET Balance = (SELECT Balance WHERE Phone = ?) + (SELECT TransferBalance FROM transfer ORDER BY CreatedAt DESC LIMIT 1)`, receiver)
			if err2 != nil {
				fmt.Println("error3", err2.Error())
				return err2
			} else {
				return nil
			}
		}
	}
}

func HistoryTopUp(db *sql.DB, phone string) ([]_entities.TopUp, error) {
	query, err := db.Query(`SELECT TopUpBalance, CreatedAt from topup WHERE TopUpPhone = ? ORDER BY CreatedAt DESC`, phone)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	var other []_entities.TopUp

	for query.Next() {
		var data _entities.TopUp
		err := query.Scan(&data.TopUpBalance, &data.CreatedAt)

		if err != nil {
			fmt.Println("error2", err.Error())
		}
		other = append(other, data)
	}
	return other, nil
}

func HistoryTransfer(db *sql.DB, phone string) ([]_entities.Transfer, error) {
	query, err := db.Query(`SELECT ReceiverPhone, TransferBalance, CreatedAt from transfer WHERE SenderPhone = ? ORDER BY CreatedAt DESC`, phone)

	if err != nil {
		fmt.Println("error1", err.Error())
	}

	var other []_entities.Transfer

	for query.Next() {
		var data _entities.Transfer
		err := query.Scan(&data.ReceiverPhone, &data.TransferBalance, &data.CreatedAt)

		if err != nil {
			fmt.Println("error2", err.Error())
		}
		other = append(other, data)
	}
	return other, nil
}
