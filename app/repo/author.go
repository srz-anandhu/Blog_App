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

func(r *Author)TableName()string{
	return " authors"
}

func (r *Author) Create(db *sql.DB) (lastInsertedID int64, err error) {

	query := `INSERT INTO` + r.TableName() + `(name)
			  VALUES ($1)
			  RETURNING id`

	if err := db.QueryRow(query,r.Name).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id due to: %w", err)
	}
	return lastInsertedID, nil
}
