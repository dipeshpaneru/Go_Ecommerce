package main

import (
	"log"
	"github.com/dipeshpaneru/Go_Ecommerce/cmd/api"
	"github.com/go-sql-driver/mysql"
	"github.com/dipeshpaneru/Go_Ecommerce/database"
	"github.com/dipeshpaneru/Go_Ecommerce/config"
	"database/sql"
)

func main() {

	sql, err := database.NewMySQLStorage(mysql.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Net:    "tcp",
		Addr:   config.Envs.DBAddress,
		DBName: config.Envs.DBName,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	initStorage(sql)

	server := api.NewAPIServer(config.Envs.Port, sql)

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