package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)


func NewPostgreSQLStorage(connStr string) (*sql.DB, error) {
	psql, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return psql, nil
}