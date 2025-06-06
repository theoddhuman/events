package model

import (
	"database/sql"
	"errors"
	"events/db"
	"events/utils"
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
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	var result sql.Result
	result, err = stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.Id = userId
	return err
}

func (user User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var password string
	err := row.Scan(&password)
	if err != nil {
		return errors.New("invalid credentials")
	}
	passwordIsValid := utils.CheckPasswordHash(user.Password, password)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
