package repo

import (
	"database/sql"
	"fmt"
	"time"
)

// Blog Model
type Blog struct {
	ID        uint16
	Title     string
	Content   string
	Status    int    // 1 - Draft, 2 - Published, 3 - Deleted
	AuthorID  uint16 // Author.ID
	UpdatedBy uint16 // User.ID
	CreatedBy uint16 // User.ID
	DeletedBy uint16 // User.ID
	DeletedAt time.Time
	Model
}

func (r *Blog) Create(db *sql.DB) (lastInsertedID int64, err error) {

	query := `INSERT INTO blogs(title,content,author_id,status,created_by)
			 VALUES ($1,$2,$3,$4,$5)`

	_, err = db.Exec(query, r.Title, r.Content, r.AuthorID, r.Status, r.CreatedBy)
	if err != nil {
		return 0, fmt.Errorf("query execution failed due to : %w", err)
	}

	query = `SELECT lastval()`
	if err := db.QueryRow(query).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id dut to : %w", err)
	}
	return lastInsertedID, nil
}
