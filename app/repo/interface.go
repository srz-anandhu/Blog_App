package repo

type Repo interface {
	Create() (lastInsertedID int64, err error)
	Update(id int) (err error)
	Delete(id int) (err error)
	GetOne(id int) (result interface{}, err error)
	GetAll() (results []interface{}, err error)
	TableName() string // Function for reuse/modify table name
}
