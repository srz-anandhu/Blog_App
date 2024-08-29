package repo

import (
	"blog/app/dto"
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

type AuthorRepo interface {
	Create(authorReq *dto.AuthorCreateRequest) (lastInsertedID int64, err error)
	Update(authorUpdateReq *dto.AuthorUpdateRequest) (err error)
	Delete(id int) (err error)
	GetOne(id int) (authorResp *dto.AuthorResponse, err error)
	GetAll() (authorsResp *[]dto.AuthorResponse, err error)
	TableName() string // Function for reuse/modify table name
}

type AuthorRepoImpl struct {
	db *sql.DB
}

// For checking implementation of AuthorRepo interface
var _ AuthorRepo = (*AuthorRepoImpl)(nil)

func NewAuthorRepo(db *sql.DB) AuthorRepo {
	return &AuthorRepoImpl{
		db: db,
	}
}

// Function for reuse/modify table name
func (r *AuthorRepoImpl) TableName() string {
	return " authors "
}

func (r *AuthorRepoImpl) Create(authorReq *dto.AuthorCreateRequest) (lastInsertedID int64, err error) {

	query := `INSERT INTO` + r.TableName() + `(name,created_by)
			  VALUES ($1,$2)
			  RETURNING id`

	if err := r.db.QueryRow(query, authorReq.Name, authorReq.CreatedBy).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id due to: %w", err)
	}
	return lastInsertedID, nil
}

func (r *AuthorRepoImpl) Update(authorUpdateReq *dto.AuthorUpdateRequest) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET name=$1,updated_at=$2,updated_by=$3
			  WHERE id=$4`

	result, err := r.db.Exec(query, authorUpdateReq.Name, time.Now().UTC(), authorUpdateReq.UpdatedBy, authorUpdateReq.ID)
	if err != nil {
		return fmt.Errorf("update query failed due to : %w", err)
	}
	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d ", authorUpdateReq.ID)
	}
	return nil
}

func (r *AuthorRepoImpl) Delete(id int) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET deleted_at=$1
		      WHERE id=$2`

	_, err = r.db.Exec(query, time.Now().UTC(), id)
	if err != nil {
		return fmt.Errorf("update query failed due to : %w", err)
	}
	return nil
}

func (r *AuthorRepoImpl) GetOne(id int) (authorResp *dto.AuthorResponse, err error) {
	query := `SELECT id,name,created_at,created_by,updated_at,updated_by,deleted_at
			  FROM` + r.TableName() +
		`WHERE id=$1`

	authorResp = &dto.AuthorResponse{}
	if err := r.db.QueryRow(query, id).Scan(&authorResp.ID, &authorResp.Name, &authorResp.CreatedAt, &authorResp.CreatedBy, &authorResp.UpdatedAt, &authorResp.UpdatedBy, &authorResp.DeletedAt); err != nil {
		return nil, fmt.Errorf("query failed due to : %w", err)
	}
	return authorResp, nil
}

func (r *AuthorRepoImpl) GetAll() (authorsResp *[]dto.AuthorResponse, err error) {
	query := `SELECT id,name,created_at,created_by,updated_at,updated_by
	         FROM ` + r.TableName() + ``

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	defer rows.Close()

	var authorsCollection []dto.AuthorResponse
	for rows.Next() {

		authors := dto.AuthorResponse{}
		if err := rows.Scan(&authors.ID, &authors.Name, &authors.CreatedAt, &authors.CreatedBy, &authors.UpdatedAt, &authors.UpdatedBy); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		authorsCollection = append(authorsCollection, authors)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return &authorsCollection, nil
}
