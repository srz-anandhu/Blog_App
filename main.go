package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
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

	// User Creation *********************************************************************************
	// Try with UNIQUE usernames..
	err := CreateUser("user1018@gmail.com", "strongpassword3")
	if err != nil {
		log.Printf("user creation failed due to: %s", err)
	} else {
		fmt.Println("User created successfully")
	}

	// Create Author *********************************************************************************
	err = CreateAuthor("author2227")
	if err != nil {
		log.Printf("author creation failed due to: %s", err)
	} else {
		fmt.Println("Author created successfully")
	}
}
// ----------------------------------------------------------------------------------------------------------------------------

func CreateUser(username, password string) error {
	// Generate salt
	salt, err := GenerateSalt(10)
	if err != nil {
		return fmt.Errorf("error generating salt: %v", err)
	}

	// Password Hashing
	hashedPassword := HashPassword(password, salt)

	query := `INSERT INTO users(username,password,salt)
			  VALUES ($1,$2,$3)`

	_, err = db.Exec(query, username, hashedPassword, salt)
	if err != nil {
		return fmt.Errorf("execution error due to : %s", err)
	}

	return nil
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

func CreateAuthor(name string) error {

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
