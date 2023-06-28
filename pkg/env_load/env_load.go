package envload

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load("/home/whoami/Documents/github/severus/auth_golang/.env")
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
	}
}
