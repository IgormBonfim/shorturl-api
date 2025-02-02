package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/igormbonfim/shorturl-api/internal/api/routes"
	"github.com/igormbonfim/shorturl-api/internal/api/services"
	"github.com/igormbonfim/shorturl-api/internal/infra/database"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env != "Production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Aviso: Não foi possível carregar .env, usando variáveis de ambiente.")
		}
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Ocorreu um erro ao inicializar o banco de dados: %v", err)
	}

	defer database.Close()

	services.RegisterServices()
	routes.RegisterRoutes()

	fmt.Println("🚀 Servidor rodando na porta 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
