package model

import (
	"database/sql"
	"events/db"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.Id = userId
	return err
}
