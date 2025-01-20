package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	envload "github.com/severusTI/auth_golang/pkg/env_load"
)

func InitDBConnection(cfg envload.Config) *sql.DB {
	dbHost := cfg.DB.Host
	dbPort := cfg.DB.Port
	dbUser := cfg.DB.User
	dbPassword := cfg.DB.Password
	dbName := cfg.DB.DBName
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("error while loading environment variables of the database")
	}

	// Monta a string de conexão
	connString := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser +
		" password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	fmt.Println(connString)
	time.Sleep(3 * time.Second)
	// Inicia a conexão com o banco de dados
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Teste de conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Database connection successful.")
	return db
}
