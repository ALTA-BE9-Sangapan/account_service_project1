package controllers

import (
	"database/sql"
	_entities"project1/entities"
)

func LoginUser(db *sql.DB, newUser _entities *project1.) (int, error) {
	
	var query = ("INSERT INTO Users (user_phone, user_id, user_name, password) VALUES (?, ?, ?, ?)")
	statement,errPrepared := db.Query(query)

	if errPrepared != nil {
		return 0,errPrepared
	}
	result, err := statement.Exec(newUser.user_phone, newUser.user_id, newUser.user_name, newUser.password)

	defer result.Close()

	if err != nil {
		return 0,err
	}else{
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}
