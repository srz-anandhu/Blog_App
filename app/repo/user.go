package repo

import (
	"blog/app/dto"
	"blog/pkg/salthash"
	"database/sql"
	"fmt"
	"time"
)

// User Model
type User struct {
	ID        int
	UserName  string
	Password  string
	Salt      string
	IsDeleted bool
	Model
	DeleteInfo
}

type UserRepo interface {
	Create(userReq *dto.UserCreateRequest) (lastInsertedID int64, err error)
	Update(userUpdateReq *dto.UserUpdateRequest) (err error)
	Delete(id int) (err error)
	GetOne(id int) (userResp *dto.UserResponse, err error)
	GetAll() (usersResp *[]dto.UserResponse, err error)
	TableName() string // Function for reuse/modify table name
}

type UserRepoImpl struct {
	db *sql.DB
}

// For checking implementation of Repo interface
var _ UserRepo = (*UserRepoImpl)(nil)

func NewUserRepo(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

// Function for reuse/modify table name
func (r *UserRepoImpl) TableName() string {
	return " users "
}

var user User

func (r *UserRepoImpl) Create(userReq *dto.UserCreateRequest) (lastInsertedID int64, err error) {
	// Generate Salt
	salt, err := salthash.GenerateSalt(10)
	if err != nil {
		return 0, fmt.Errorf("error generating salt : %w", err)
	}
	// Hash Password
	PasswordString := salthash.HashPassword(userReq.Password, salt)

	query := `INSERT INTO` + r.TableName() + `(username,password,salt)
			  VALUES ($1,$2,$3)
			  RETURNING id`

	if err := r.db.QueryRow(query, userReq.UserName, PasswordString, salt).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("query execution failed due to : %w", err)
	}

	return lastInsertedID, nil
}

func (r *UserRepoImpl) Update(userUpdateReq *dto.UserUpdateRequest) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET username=$1,password=$2,updated_at=$3
			  WHERE id=$4`

	passwordHash := salthash.HashPassword(userUpdateReq.Password, user.Salt)

	result, err := r.db.Exec(query, userUpdateReq.Username, passwordHash, time.Now().UTC(), userUpdateReq.ID)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}

	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d", userUpdateReq.ID)
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

func (r *UserRepoImpl) GetAll() (usersResp *[]dto.UserResponse, err error) {
	query := `SELECT id,username,password,salt,created_at,updated_at,is_deleted,deleted_at
	        FROM` + r.TableName() + ``

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	defer rows.Close()
	var userCollection []dto.UserResponse

	for rows.Next() {
		var user dto.UserResponse
		if err := rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt, &user.IsDeleted, &user.DeletedAt); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		userCollection = append(userCollection, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return &userCollection, nil
}

func (r *UserRepoImpl) GetOne(id int) (userResp *dto.UserResponse, err error) {
	query := `SELECT id,username,password,salt,created_at,updated_at,is_deleted,deleted_at
			  FROM` + r.TableName() +
		`WHERE id=$1`
	userResp = &dto.UserResponse{}
	if err := r.db.QueryRow(query, id).Scan(&userResp.ID, &userResp.UserName, &userResp.Password, &userResp.Salt, &userResp.CreatedAt, &userResp.UpdatedAt, &userResp.IsDeleted, &userResp.DeletedAt); err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	return userResp, nil
}
