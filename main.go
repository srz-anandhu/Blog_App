package main

import (
	"blog/app"
	"blog/app/db"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//******************************************************
	
	// DB initialization
	db, err := db.InitDB()
	if err != nil {
		log.Printf("db connection error due to : %s", err)
	}
	app.Start(db)
	// close db connection before main function exit
	defer db.Close()

	// log.Fatal(http.ListenAndServe(":8080", r))

	// Get All Blogs

	// Get One Blog

	// Create Blog
	// r.Post("/blogs", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("blog created"))
	// })

	// // Update Blog
	// r.Put("/blog/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("blog updated"))
	// })

	// // Delete Blog
	// r.Delete("/blog/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("blog deleted "))
	// })

	// // Get All User
	// r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("got all users"))
	// })

	// // Get One User
	// r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("got a single user"))
	// })

	// // Create User
	// r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("user created"))
	// })

	// // Update User
	// r.Put("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("user updated"))
	// })

	// // Delete User
	// r.Delete("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("deleted user"))
	// })

	// // Get All Authors
	// r.Get("/authors", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("got all authors"))
	// })

	// // Get One Author
	// r.Get("/author/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("got one author"))
	// })

	// // Create Author
	// r.Post("/author", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("created author"))
	// })

	// // Update Author
	// r.Put("/author/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("author updated"))
	// })

	// // Delete Author
	// r.Delete("/author/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("author deleted"))
	// })

	//******************************************************

	//var user repo.User // creating an instance of user
	//var author repo.Author // creating an instance of author
	//var blog repo.Blog // creating an instance of blog

	// User Creation
	// user.UserName = "user7010@gmail.com"
	// user.Password = "strongPassword"

	// userID, err := user.Create(db)
	// if err != nil {
	// 	log.Printf("user creation failed due to : %s", err)
	// } else {
	// 	fmt.Printf("user created with ID: %d", userID)
	// }
	// user.ID = 10
	// if err := user.Delete(db); err != nil {
	// 	log.Printf("can't delete user with ID: %d", user.ID)
	// } else {
	// 	fmt.Println("deleted user successfully")
	// }
	//	Get all users
	// users, err := user.GetAll(db)
	// if err != nil {
	// 	log.Printf("can't get users due to : %s", err)
	// } else {
	// 	for _, user := range users {
	// 		u, ok := user.(repo.User)
	// 		if !ok {
	// 			log.Println("type assertion failed")
	// 		}
	// 		fmt.Printf("\n ID: %d \n UserName: %s \n Password: %s \n Created AT: %s \n Updated AT: %s", u.ID, u.UserName, u.Password, u.CreatedAt, u.UpdatedAt)
	// 	}

	// Get one user
	// user.ID = 20
	// singleUser,err:=user.GetOne(db)
	// if err!=nil{
	// 	log.Printf("can't get user due to : %s",err)
	// }else{
	// 	// type assertion to convert singleUser to User
	// 	u, ok:= singleUser.(repo.User)
	// 	if !ok{
	// 		log.Println("type assertion failed")
	// 	}
	// 	fmt.Printf("\n ID: %d \n UserName: %s \n Password: %s \n Created AT: %s \n Updated AT: %s",u.ID, u.UserName, u.Password, u.CreatedAt, u.UpdatedAt)
	// }

	//Author Creation
	// author.Name = "author1111"

	// authorID, err := author.Create(db)
	// if err != nil {
	// 	log.Printf("author creation failed due to : %s", err)
	// } else {
	// 	fmt.Printf("author created with ID: %d", authorID)
	// }

	// Author Updation
	// author.ID=7
	// author.Name = "anandhu"
	// if err:=author.Update(db);err!=nil{
	// 	log.Printf("cant't update author due to : %s",err)
	// }else{
	// 	fmt.Printf("updated author with ID: %d",author.ID)
	// }

	//Blog Creation
	// blog.Title = "Blog Title44444"
	// blog.Content = "Content of blog 44444"
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

	// // Delete Blog
	// blog.DeletedBy = 20 // User.ID
	// blog.ID = 6
	// if err := blog.Delete(db); err != nil {
	// 	log.Printf("blog deletion failed due to : %s", err)
	// } else {
	// 	log.Printf("blog with ID: %d deleted successfully", blog.ID)
	// }

	// Get Single Blog
	// blog.ID = 8
	// singleBlog, err := blog.GetOne(db)
	// if err != nil {
	// 	log.Printf("Can't get blog due to : %s", err)
	// } else {
	// 	// Type assertion to convert singleBlog to Blog
	// 	b,ok :=singleBlog.(repo.Blog)
	// 	if !ok{
	// 		log.Println("type assertion to blog failed ")
	// 	}
	// 	fmt.Printf("\n ID: %d \n Title: %s \n Content: %s \n Author: %d \n Created At: %s \n Updated At: %s",b.ID, b.Title, b.Content, b. AuthorID, b.CreatedAt, b.UpdatedAt)
	// }

	// // Get All Blogs
	// blogs, err := blog.GetAll(db)
	// if err != nil {
	// 	log.Printf("can't get blogs due to : %s", err)
	// }else{

	// 	for _, blog := range blogs {
	// 		// Type assertion to convert blogs to Blog
	// 		b, ok:= blog.(repo.Blog)
	// 		if !ok{
	// 			log.Println("type assertion failed")
	// 		}
	// 		fmt.Printf("\n ID: %d \n Title: %s \n Content: %s \n Author: %d \n Created At %s \n Updated At: %s",b.ID, b.Title, b.Content, b.AuthorID, b.CreatedAt, b.UpdatedAt)
	// 	}
	// }
}
