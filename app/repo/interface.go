package repo

import "database/sql"

type Repo interface {
	Create(db *sql.DB) (lastInsertedID int64, err error)
	Update(db *sql.DB) (err error)
	Delete(db *sql.DB) (err error)
	GetOne(db *sql.DB) (blog Blog, err error)
	GetAll(db *sql.DB) (blogs []Blog, err error)
}
