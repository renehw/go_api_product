package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
