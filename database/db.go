package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)


func NewPostgreSQLStorage(connStr string) (*sql.DB, error) {
	log.Printf("Connecting to PostgreSQL with connection string: %s", connStr)
	psql, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return psql, nil
}