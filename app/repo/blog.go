package repo

import (
	"blog/app/dto"
	"database/sql"
	"fmt"
	"time"
)

// Blog Model
type Blog struct {
	ID       int
	Title    string
	Content  string
	Status   int // 1 - Draft, 2 - Published, 3 - Deleted
	AuthorID int // Author.ID
	Model
	DeleteInfo
}

type BlogRepo interface {
	Create(blogReq *dto.BlogCreateRequest) (lastInsertedID int64, err error)
	Update(blogUpdateReq *dto.BlogUpdateRequest) (err error)
	Delete(id int) (err error)
	GetOne(id int) (blogResp *dto.BlogResponse, err error)
	GetAll() (blogsResp *[]dto.BlogResponse, err error)
	TableName() string // Function for reuse/modify table name
}

type BlogRepoImpl struct {
	db *sql.DB
}

// For checking implementation of Repo interface
var _ BlogRepo = (*BlogRepoImpl)(nil)

func NewBlogRepo(db *sql.DB) BlogRepo {
	return &BlogRepoImpl{
		db: db,
	}
}

// Function for reuse/modify table name
func (r *BlogRepoImpl) TableName() string {
	return " blogs "
}

func (r *BlogRepoImpl) Create(blogReq *dto.BlogCreateRequest) (lastInsertedID int64, err error) {

	query := `INSERT INTO` + r.TableName() + `(title,content,author_id,status,created_by) 
			 VALUES ($1,$2,$3,$4,$5)
			 RETURNING id`

	if err := r.db.QueryRow(query, blogReq.Title, blogReq.Content, blogReq.AuthorID, blogReq.Status, blogReq.CreatedBy).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id dut to : %w", err)
	}
	return lastInsertedID, nil
}

func (r *BlogRepoImpl) Update(blogUpdateReq *dto.BlogUpdateRequest) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET title=$1,content=$2,status=$3,updated_at=$4,updated_by=$5
	    WHERE id=$6
	    AND status
	    IN (1,2)`

	result, err := r.db.Exec(query, blogUpdateReq.Title, blogUpdateReq.Content, blogUpdateReq.Status, time.Now().UTC(), blogUpdateReq.UpdatedBy, blogUpdateReq.ID)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}
	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to: %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no blogs with id=%d or status in 1 or 2", blogUpdateReq.ID)
	}

	return nil
}

// Soft Delete
func (r *BlogRepoImpl) Delete(id int) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET deleted_by=$1,deleted_at=$2,status=$3
	WHERE id=$4`
	var blog Blog
	_, err = r.db.Exec(query, blog.DeletedBy, time.Now().UTC(), 3, id)
	if err != nil {
		return fmt.Errorf("delete query execution failed due to: %w", err)
	}
	return nil
}

func (r *BlogRepoImpl) GetOne(id int) (blogResp *dto.BlogResponse, err error) {
	query := `SELECT id,title,content,author_id,status,created_at,created_by,updated_at,updated_by,deleted_at,deleted_by FROM` + r.TableName() + `WHERE id=$1`
	blogResp = &dto.BlogResponse{}
	if err := r.db.QueryRow(query, id).Scan(&blogResp.ID, &blogResp.Title, &blogResp.Content, &blogResp.AuthorID, &blogResp.Status, &blogResp.CreatedAt, &blogResp.CreatedBy, &blogResp.UpdatedAt, &blogResp.UpdatedBy, &blogResp.DeletedAt, &blogResp.DeletedBy); err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	return blogResp, nil
}

func (r *BlogRepoImpl) GetAll() (blogsResp *[]dto.BlogResponse, err error) {
	query := `SELECT id,title,content,author_id,status,created_at,created_by,updated_at,updated_by,deleted_at,deleted_by
			 FROM` + r.TableName() + `` // blogs

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %s", err)
	}
	defer rows.Close()
	var blogsCollection []dto.BlogResponse
	for rows.Next() {

		blogs := dto.BlogResponse{}
		if err := rows.Scan(&blogs.ID, &blogs.Title, &blogs.Content, &blogs.AuthorID, &blogs.Status, &blogs.CreatedAt, &blogs.CreatedBy, &blogs.UpdatedAt, &blogs.UpdatedBy, &blogs.DeletedAt, &blogs.DeletedBy); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		blogsCollection = append(blogsCollection, blogs)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return &blogsCollection, nil
}
