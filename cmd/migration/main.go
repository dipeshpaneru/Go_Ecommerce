package main

import (
	_ "github.com/lib/pq"
	"github.com/dipeshpaneru/Go_Ecommerce/database"
	"github.com/dipeshpaneru/Go_Ecommerce/config"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	sql, err := database.NewPostgreSQLStorage(
		"host=" + config.Envs.Host +
		" port=" + config.Envs.Port +
		" user=" + config.Envs.DBUser +
		" dbname=" + config.Envs.DBName +
		" sslmode=disable",
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}


	driver, err := postgres.WithInstance(sql, &postgres.Config{})

	if err != nil {
		log.Fatalf("Failed to create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migration/migrations",
		"postgres",
		 driver)
	
	if err != nil {
		log.Fatalf("Failed to create database driver: %v", err)
	}	
	
	m.up hello helllo


}
