package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	user     = "postgres"
	password = "password"
	host     = "localhost"
	port     = 5432
	dbname   = "blogdatabase"
)

func InitDB() (*sql.DB, error) {

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode = disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", connectionString)
	// DSN parse error or initialization error
	if err != nil {
		log.Fatal(err)
	}
	// connection checking
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connected to database....")
	return db, nil
}
