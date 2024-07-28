package main

import (
	"blog/app/repo"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "password"
	host     = "localhost"
	port     = 5432
	dbname   = "blogdatabase"
)

var db *sql.DB
var err error

func main() {

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode = disable", user, password, host, port, dbname)

	db, err = sql.Open("postgres", connectionString)
	// DSN parse error or initialization error
	if err != nil {
		log.Fatal(err)
	}
	// close db connection before main function exit
	defer db.Close()
	// connection checking
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connected to database....")

	// User Creation
	// var user repo.User // creating an instance of user
	// user.UserName = "user708@gmail.com"
	// user.Password = "strongPassword"

	// userID, err := user.Create(db)
	// if err != nil {
	// 	log.Printf("user creation failed due to : %s", err)
	// } else {
	// 	fmt.Printf("user created with ID: %d", userID)
	// }

	// Author Creation
	var author repo.Author
	author.Name = "author1111"

	authorID, err := author.Create(db)
	if err != nil {
		log.Printf("author creation failed due to : %s", err)
	} else {
		fmt.Printf("author created with ID: %d", authorID)
	}

	// create Blog ************************************************************************************
	// title,content,authrorId,status=(1,2,3:=drafted,published,deleted),userId
	// err = createBlog("blog10", "This is content for blog10", 3, 2, 20)
	// if err != nil {
	// 	log.Printf("blog creation failed due to : %s", err)
	// } else {
	// 	log.Println("blog created successfully")
	// }

	// delete Blog ************************************************************************************
	// blogid,userid
	// err = deleteBlog(3, 12)
	// if err != nil {
	// 	log.Printf("blog deletion failed due to : %s", err)
	// } else {
	// 	log.Println("blog deleted successfully")
	// }

	// read Blog *************************************************************************************
	// title,content,authorId
	// title, content, authorId, created_at, updated_at, err := readBlog(6)
	// if err != nil {
	// 	log.Printf("can't get blog due to : %s", err)
	// } else {
	// 	fmt.Printf("Title: %s \n Content: %s \n AuthorId: %d\n Created At:%s\n Updated At: %s\n ", title, content, authorId, created_at, updated_at)
	// }

	// update Blog ***********************************************************************************
	// title,content,blogId
	// err = updateBlog("updated title", "this is updated content ", 4)
	// if err != nil {
	// 	log.Printf("blog updation failed due to : %s", err)
	// } else {
	// 	fmt.Println("blog updated successfully")
	// }

	// read all Blogs *********************************************************************************
	// blogs, err := readAllBlogs()
	// if err != nil {
	// 	log.Printf("getting blogs failed due to : %s", err)
	// } else {
	// 	for _, blog := range blogs {
	// 		fmt.Printf("BlogID: %d \n Title: %s \n Content:%s \n AuthorId: %d \n Created At: %s \n Updated At: %s \n", blog.id, blog.title, blog.content, blog.authorId, blog.created_at, blog.updated_at)
	// 	}
	// }
}

// ----------------------------------------------------------------------------------------------------------------------------

// func createBlog(title, content string, authorId, status, userId uint16) error {
// 	query := `INSERT INTO blogs(title,content,author_id,status,created_by)
// 			 VALUES ($1,$2,$3,$4,$5)`

// 	_, err = db.Exec(query, title, content, authorId, status, userId)
// 	if err != nil {
// 		return fmt.Errorf("query execution failed due to : %w", err)
// 	}
// 	return nil
// }

// // soft delete
// func deleteBlog(id, userId uint16) error {
// 	query := `UPDATE blogs
// 			 SET deleted_by=$1,deleted_at=$2,status=$3
// 			 WHERE id=$4`

// 	_, err = db.Exec(query, userId, time.Now().UTC(), 3, id)
// 	if err != nil {
// 		return fmt.Errorf("delete query execution failed due to: %w", err)
// 	}
// 	return nil
// }

// func readBlog(id uint16) (string, string, uint16, time.Time, time.Time, error) {
// 	var (
// 		title      string
// 		content    string
// 		authorId   uint16
// 		created_at time.Time
// 		updated_at time.Time
// 	)

// 	query := `SELECT title,content,author_id,created_at,updated_at FROM blogs WHERE id=$1 AND status = 2`

// 	if err := db.QueryRow(query, id).Scan(&title, &content, &authorId, &created_at, &updated_at); err != nil {
// 		return "", "", 0, time.Time{}, time.Time{}, err
// 	}
// 	return title, content, authorId, created_at, updated_at, nil
// }

// func updateBlog(title, content string, id uint16) error {
// 	query := `UPDATE blogs
// 			  SET title=$1,content=$2,updated_at=$3
// 			  WHERE id=$4
// 			  AND status
// 			  IN (1,2)
// 			 `

// 	result, err := db.Exec(query, title, content, time.Now().UTC(), id)
// 	if err != nil {
// 		return fmt.Errorf("query execution failed due to : %w", err)
// 	}
// 	isAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return fmt.Errorf("no affected rows due to: %w", err)
// 	}
// 	if isAffected == 0 {
// 		return fmt.Errorf("no blogs with id=%d or status in 1 or 2", id)
// 	}

// 	return nil
// }

// type blog struct {
// 	id         uint16
// 	title      string
// 	authorId   uint16
// 	content    string
// 	created_at time.Time
// 	updated_at time.Time
// }

// func readAllBlogs() ([]blog, error) {

// 	query := `SELECT id,title,author_id,content,created_at,updated_at
// 			 FROM blogs`

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("query execution failed : %w", err)
// 	}

// 	defer rows.Close()

// 	var blogs []blog
// 	for rows.Next() {
// 		var blog blog
// 		if err := rows.Scan(&blog.id, &blog.title, &blog.authorId, &blog.content, &blog.created_at, &blog.updated_at); err != nil {
// 			return nil, fmt.Errorf("row scan failed due to : %w", err)
// 		}
// 		blogs = append(blogs, blog)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("row iteration failed due to : %w", err)
// 	}
// 	return blogs, nil
// }
