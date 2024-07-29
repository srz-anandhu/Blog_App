package repo

import (
	"blog/pkg/salthash"
	"database/sql"
	"fmt"
	"time"
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

func (r *User) Create(db *sql.DB) (lastInsertedID int64, err error) {
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
	query = `SELECT lastval()`
	if err := db.QueryRow(query).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id using query due to : %w", err)
	}
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, fmt.Errorf("couldn't get last inserted id due to : %w", err)
	// }
	return lastInsertedID, nil
}

func (r *User) Update(db *sql.DB) (err error) {
	query := `UPDATE users
			 SET username=$1,password=$2,updated_at=$3
			 WHERE id=$4`

	result, err := db.Exec(query, r.UserName, r.Password, time.Now().UTC(), r.ID)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}

	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d", r.ID)
	}
	return nil
}

// Soft Delete
func (r *User) Delete(db *sql.DB) (err error) {
	query := `UPDATE users
			 SET is_deleted=$1
			 WHERE id=$2`

	_, err = db.Exec(query, true, r.ID)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}
	return nil
}
