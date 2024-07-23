package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"

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

	// User Creation *********************************************************************************
	// Try with UNIQUE usernames..
	// err := createUser("user1018@gmail.com", "strongpassword3")
	// if err != nil {
	// 	log.Printf("user creation failed due to: %s", err)
	// } else {
	// 	fmt.Println("User created successfully")
	// }

	// Create Author *********************************************************************************
	// err = createAuthor("author2227")
	// if err != nil {
	// 	log.Printf("author creation failed due to: %s", err)
	// } else {
	// 	fmt.Println("Author created successfully")
	// }

	// create Blog ************************************************************************************
	// title,content,authrorId,status=(1,2,3:=drafted,published,deleted),userId
	// err = createBlog("blog4", "This is content for blog4", 3, 2, 5)
	// if err != nil {
	// 	log.Printf("blog creation failed due to : %s", err)
	// } else {
	// 	log.Println("blog created successfully")
	// }

	// delete Blog ************************************************************************************
	// blogid,userid
	// err = deleteBlog(4, 5)
	// if err != nil {
	// 	log.Printf("blog deletion failed due to : %s", err)
	// } else {
	// 	log.Println("blog deleted successfully")
	// }

	// read Blog *************************************************************************************
	// title,content,authorId
	title, content, authorId, created_at, updated_at, err := readBlog(6)
	if err != nil {
		log.Printf("can't get blog due to : %s", err)
	} else {
		fmt.Printf("Title: %s \n Content: %s \n AuthorId: %d\n Created At:%s\n Updated At: %s\n ", title, content, authorId, created_at, updated_at)
	}
}

// ----------------------------------------------------------------------------------------------------------------------------

func createUser(username, password string) error {
	// Generate salt
	salt, err := generateSalt(10)
	if err != nil {
		return fmt.Errorf("error generating salt: %v", err)
	}

	// Password Hashing
	hashedPassword := hashPassword(password, salt)

	query := `INSERT INTO users(username,password,salt)
			  VALUES ($1,$2,$3)`

	_, err = db.Exec(query, username, hashedPassword, salt)
	if err != nil {
		return fmt.Errorf("execution error due to : %s", err)
	}

	return nil
}

// Generate a random salt of the given length
func generateSalt(length uint8) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	saltString := hex.EncodeToString(bytes)
	return saltString, nil
}

// Hashes the password with given salt (SHA-256)
func hashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashBytes := hash.Sum(nil)
	hashedPass := hex.EncodeToString(hashBytes)
	return hashedPass
}

func createAuthor(name string) error {

	query := `INSERT INTO authors(name)
			  VALUES ($1)`

	// if err := db.QueryRow(query, name, time.Now(), time.Now(), true).Scan(&authorId); err != nil {
	// 	return 0, err
	// }
	_, err = db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("execution error due to: %s ", err)
	}
	return nil
}

func createBlog(title, content string, authorId, status, userId uint16) error {
	query := `INSERT INTO blogs(title,content,author_id,status,created_by)
			 VALUES ($1,$2,$3,$4,$5)`

	_, err = db.Exec(query, title, content, authorId, status, userId)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %s", err)
	}
	return nil
}

// soft delete
func deleteBlog(id, userId uint16) error {
	query := `UPDATE blogs
			 SET deleted_by=$1,deleted_at=$2,status=$3
			 WHERE id=$4`

	_, err = db.Exec(query, userId, time.Now().UTC(), 3, id)
	if err != nil {
		return fmt.Errorf("delete query execution failed due to: %s", err)
	}
	return nil
}

func readBlog(id uint16) (string, string, uint16, time.Time, time.Time, error) {
	var title string
	var content string
	var authorId uint16
	var created_at time.Time
	var updated_at time.Time

	query := `SELECT title,content,author_id,created_at,updated_at FROM blogs WHERE id=$1 AND status = 2`

	if err := db.QueryRow(query, id).Scan(&title, &content, &authorId, &created_at, &updated_at); err != nil {
		return "", "", 0, time.Time{}, time.Time{}, err
	}
	return title, content, authorId, created_at, updated_at, nil
}
