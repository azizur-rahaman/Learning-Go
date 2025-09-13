package models

import (
	"azizur/rest-api/db"
	"azizur/rest-api/util"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES(?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	hash, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hash)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := util.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
