package repo

import (
	"blog/pkg/salthash"
	"database/sql"
	"fmt"
)

// User Model
type User struct {
	ID        uint16
	UserName  string
	Password  string
	Salt      string
	IsDeleted bool
	Model
}


func (r *User)Create(db *sql.DB) (lastInsertedID int64, err error) {
	// Generate Salt
	salt, err := salthash.GenerateSalt(100)
	if err != nil {
		return 0, fmt.Errorf("error generating salt : %w", err)
	}
	// Hash Password
	PasswordString := salthash.HashPassword(r.Password, salt)

	query := `INSERT INTO users(username,password,salt)
			  VALUES ($1,$2,$3)`

	_, err = db.Exec(query, r.UserName, PasswordString, salt)
	if err != nil {
		return 0, fmt.Errorf("query execution failed due to : %w", err)
	}
	query =`SELECT lastval()`
	if err:=db.QueryRow(query).Scan(&lastInsertedID);err !=nil{
		return 0,fmt.Errorf("couldn't get last inserted id using query due to : %w",err)
	}
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, fmt.Errorf("couldn't get last inserted id due to : %w", err)
	// }
	return lastInsertedID, nil
}

