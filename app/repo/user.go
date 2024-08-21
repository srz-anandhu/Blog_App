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
	DeleteInfo
}

type UserRepoImpl struct {
	db *sql.DB
}

var _ Repo = (*UserRepoImpl)(nil)

func NewUserRepo(db *sql.DB) Repo {
	return &UserRepoImpl{
		db: db,
	}
}

func (r *UserRepoImpl) TableName() string {
	return " users "
}

var user User

func (r *UserRepoImpl) Create() (lastInsertedID int64, err error) {
	// Generate Salt
	salt, err := salthash.GenerateSalt(10)
	if err != nil {
		return 0, fmt.Errorf("error generating salt : %w", err)
	}
	// Hash Password
	PasswordString := salthash.HashPassword(user.Password, salt)

	query := `INSERT INTO` + r.TableName() + `(username,password,salt)
			  VALUES ($1,$2,$3)
			  RETURNING id`

	if err := r.db.QueryRow(query, user.UserName, PasswordString, salt).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("query execution failed due to : %w", err)
	}

	return lastInsertedID, nil
}

func (r *UserRepoImpl) Update(id int) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET username=$1,password=$2,updated_at=$3,updated_by=$4
			  WHERE id=$5`

	passwordHash := salthash.HashPassword(user.Password, user.Salt)

	result, err := r.db.Exec(query, user.UserName, passwordHash, time.Now().UTC(), user.UpdatedBy, id)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}

	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d", id)
	}
	return nil
}

// Soft Delete
func (r *UserRepoImpl) Delete(id int) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET is_deleted=$1,deleted_at=$2
			  WHERE id=$3`

	_, err = r.db.Exec(query, true, time.Now().UTC(), id)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}
	return nil
}

func (r *UserRepoImpl) GetAll() (results []interface{}, err error) {
	query := `SELECT id,username,password,created_at,updated_at 
	        FROM` + r.TableName() + `WHERE is_deleted=false`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		// var user User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		results = append(results, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return results, nil
}

func (r *UserRepoImpl) GetOne(id int) (result interface{}, err error) {
	query := `SELECT id,username,password,created_at,updated_at
			  FROM` + r.TableName() +
		`WHERE id=$1
			  AND 
			  is_deleted=false`
	// var user User
	if err := r.db.QueryRow(query, id).Scan(&user.ID, &user.UserName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	return user, nil
}
