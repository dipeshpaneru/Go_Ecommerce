package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)


func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	sql, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	return sql, nil
}