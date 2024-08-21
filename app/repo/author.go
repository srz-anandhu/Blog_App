package repo

import (
	"database/sql"
	"fmt"
	"time"
)

// Author Model
type Author struct {
	ID   int
	Name string
	Model
	DeleteInfo
}

type AuthorRepoImpl struct {
	db *sql.DB
}

// For checking implementation of Repo interface
var _ Repo = (*AuthorRepoImpl)(nil)

func NewAuthorRepo(db *sql.DB) Repo {
	return &AuthorRepoImpl{
		db: db,
	}
}

func (r *AuthorRepoImpl) TableName() string {
	return " authors "
}

var author Author

func (r *AuthorRepoImpl) Create() (lastInsertedID int64, err error) {

	query := `INSERT INTO` + r.TableName() + `(name)
			  VALUES ($1)
			  RETURNING id`

	if err := r.db.QueryRow(query, author.Name).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id due to: %w", err)
	}
	return lastInsertedID, nil
}

func (r *AuthorRepoImpl) Update(id int) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET name=$1,updated_at=$2,updated_by=$3
			  WHERE id=$4`

	result, err := r.db.Exec(query, author.Name, time.Now().UTC(), author.UpdatedBy, id)
	if err != nil {
		return fmt.Errorf("update query failed due to : %w", err)
	}
	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d ", author.ID)
	}
	return nil
}

func (r *AuthorRepoImpl) Delete(id int) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET deleted_at=$1,deleted_by=$2
		      WHERE id=$3`

	_, err = r.db.Exec(query, time.Now().UTC(), author.DeletedBy, id)
	if err != nil {
		return fmt.Errorf("update query failed due to : %w", err)
	}
	return nil
}

func (r *AuthorRepoImpl) GetOne(id int) (result interface{}, err error) {
	query := `SELECT id,name,created_at,updated_at
			  FROM` + r.TableName() +
		`WHERE id=$1`

	if err := r.db.QueryRow(query, id).Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt); err != nil {
		return nil, fmt.Errorf("query failed due to : %w", err)
	}
	return author, nil
}

func (r *AuthorRepoImpl) GetAll() (results []interface{}, err error) {
	query := `SELECT id,name,created_at,updated_at
	         FROM ` + r.TableName() + ``

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		// var author Author
		if err := rows.Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		results = append(results, author)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return results, nil
}
