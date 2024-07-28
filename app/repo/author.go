package repo

import (
	"database/sql"
	"fmt"
)

// Author Model
type Author struct {
	ID   uint16
	Name string
	Model
}

func (r *Author) Create(db *sql.DB) (lastInsertedID int64, err error) {

	query := `INSERT INTO authors(name)
			  VALUES ($1)`

	_, err = db.Exec(query, r.Name)
	if err != nil {
		return 0, fmt.Errorf("execution error due to: %w ", err)
	}

	query = `SELECT lastval()`

	if err := db.QueryRow(query).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id due to: %w", err)
	}
	return lastInsertedID, nil
}
