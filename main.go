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

	//var user repo.User     // creating an instance of user
	//var author repo.Author // creating an instance of author
	var blog repo.Blog // creating an instance of blog

	// // User Creation
	// user.UserName = "user708@gmail.com"
	// user.Password = "strongPassword"

	// userID, err := user.Create(db)
	// if err != nil {
	// 	log.Printf("user creation failed due to : %s", err)
	// } else {
	// 	fmt.Printf("user created with ID: %d", userID)
	// }

	// //Author Creation
	// author.Name = "author1111"

	// authorID, err := author.Create(db)
	// if err != nil {
	// 	log.Printf("author creation failed due to : %s", err)
	// } else {
	// 	fmt.Printf("author created with ID: %d", authorID)
	// }

	//Blog Creation
	// blog.Title = "Blog Title2222"
	// blog.Content = "Content of blog 2222"
	// blog.AuthorID = 8
	// blog.Status = 2 // Published
	// blog.CreatedBy = 20

	// blogID, err := blog.Create(db)
	// if err != nil {
	// 	log.Printf("blog creation failed due to : %s", err)
	// } else {
	// 	fmt.Printf("blog created successfully with ID : %d", blogID)
	// }

	// Blog Updation
	// blog.Title = "Blog Title Updated to 3333"
	// blog.Content = "Content of Blog Updated to 3333"
	// blog.ID = 9
	// blog.UpdatedBy = 33
	// if err = blog.Update(db); err != nil {
	// 	log.Printf("blog updation failed due to : %s", err)
	// } else {
	// 	fmt.Printf("blog with ID: %d updated successfully", blog.ID)
	// }

	// Delete Blog
	// blog.DeletedBy = 20 // User.ID
	// blog.ID = 6
	// if err := blog.Delete(db); err != nil {
	// 	log.Printf("blog deletion failed due to : %s", err)
	// } else {
	// 	log.Printf("blog with ID: %d deleted successfully", blog.ID)
	// }

	// Get Single Blog
	blog.ID = 8
	singleBlog,err:=blog.GetOne(db)
	if err!=nil{
		log.Printf("Can't get blog due to : %s",err)
	}else{
		fmt.Printf("Title: %s \n Content: %s \n Author: %d \n Created At: %s \n Updated At: %s", singleBlog.Title,singleBlog.Content,singleBlog.AuthorID,singleBlog.CreatedAt,singleBlog.UpdatedAt)
	}

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
