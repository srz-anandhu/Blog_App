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

var Db *sql.DB
var err error

func main() {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode = disable", user, password, host, port, dbname)

	Db, err = sql.Open("postgres", connectionString)
	// DSN parse error or initialization error
	if err != nil {
		log.Fatal(err)
	}
	// close db connection before main function exit
	defer Db.Close()
	// connection checking
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connected to database....")

	// User Creation *********************************************************************************
	// Try with UNIQUE usernames..
	userId, err := CreateUser("anandhu123@gmail.com", "strongpassword3")
	if err != nil {
		log.Printf("user creation failed due to: %s", err)
	} else {
		fmt.Printf("User created with ID: %d", userId)
	}

	// Create Author *********************************************************************************
	authorId, err := CreateAuthor("author1111")
	if err != nil {
		log.Printf("author creation failed due to: %s", err)
	} else {
		fmt.Printf("Author created with ID: %d", authorId)
	}
}

func CreateUser(username, password string) (uint16, error) {
	// Generate salt
	salt, err := GenerateSalt(10)
	if err != nil {
		log.Printf("error generating salt: %v", err)
	}

	// Password Hashing
	hashedPassword := HashPassword(password, salt)

	var userId uint16
	query := `INSERT INTO users(username,password,salt,created_at,updated_at,is_deleted)
			VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	if err := Db.QueryRow(query, username, hashedPassword, salt, time.Now(), time.Now(), false).Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}

// Generate a random salt of the given length
func GenerateSalt(length uint8) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	saltString := hex.EncodeToString(bytes)
	return saltString, nil
}

// Hashes the password with given salt (SHA-256)
func HashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashBytes := hash.Sum(nil)
	hashedPass := hex.EncodeToString(hashBytes)
	return hashedPass
}

func CreateAuthor(name string) (uint16, error) {
	var authorId uint16

	query := `INSERT INTO authors(name,created_at,updated_at,status)
			VALUES ($1,$2,$3,$4)
			RETURNING id`

	if err := Db.QueryRow(query, name, time.Now(), time.Now(), true).Scan(&authorId); err != nil {
		return 0, err
	}
	return authorId, nil
}
