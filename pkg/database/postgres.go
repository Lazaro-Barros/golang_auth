package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func InitDBConnection() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
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

func InitDBTESTConnection() *sql.DB {
	dbTestHost := os.Getenv("DB_TEST_HOST")
	dbTestPort := os.Getenv("DB_TEST_PORT")
	dbTestUser := os.Getenv("DB_TEST_USER")
	dbTestPassword := os.Getenv("DB_TEST_PASSWORD")
	dbTestName := os.Getenv("DB_TEST_NAME")
	if dbTestHost == "" || dbTestPort == "" || dbTestUser == "" || dbTestPassword == "" || dbTestName == "" {
		log.Fatal("error while loading environment variables of the database")
	}

	// Monta a string de conexão
	connString := "host=" + dbTestHost + " port=" + dbTestPort + " user=" + dbTestUser +
		" password=" + dbTestPassword + " dbname=" + dbTestName + " sslmode=disable"

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

	if err = ClearTESTDatabase(db); err != nil {
		log.Fatalf("Failed to clear test database: %v", err)
	}

	log.Println("TEST Database connection successful.")
	return db
}

func ClearTESTDatabase(db *sql.DB) error {
	// Tabelas a serem limpas
	tables := []string{"users"}

	// Executa uma transação para garantir atomicidade das operações
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Limpa cada tabela individualmente
	for _, table := range tables {
		_, err := tx.Exec("DELETE FROM " + table)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	// Efetiva as alterações realizadas
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
