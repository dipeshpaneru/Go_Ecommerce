package main

import (
	"log"
	"github.com/dipeshpaneru/Go_Ecommerce/cmd/api"
	_ "github.com/lib/pq"
	"github.com/dipeshpaneru/Go_Ecommerce/database"
	"github.com/dipeshpaneru/Go_Ecommerce/config"
	"database/sql"
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

	initStorage(sql)

	server := api.NewAPIServer(":8080", sql)

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to the database")
}