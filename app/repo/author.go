package repo

import (
	"database/sql"
	"fmt"
	"time"
)

// Author Model
type Author struct {
	ID   uint16
	Name string
	Model
}

//var _ Repo = (*Author)(nil)

func (r *Author) TableName() string {
	return " authors "
}

var author Author

func (r *Author) Create(db *sql.DB) (lastInsertedID int64, err error) {

	query := `INSERT INTO` + r.TableName() + `(name)
			  VALUES ($1)
			  RETURNING id`

	if err := db.QueryRow(query, r.Name).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id due to: %w", err)
	}
	return lastInsertedID, nil
}

func (r *Author) Update(db *sql.DB) (err error) {
	query := `UPDATE` + r.TableName() +
		     `SET name=$1,updated_at=$2
			  WHERE id=$3`

	result, err := db.Exec(query, r.Name, time.Now().UTC(), r.ID)
	if err != nil {
		return fmt.Errorf("update query failed due to : %w", err)
	}
	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d ", r.ID)
	}
	return nil
}
