package main

import (
	"log"

	handlers "github.com/96vini/n0tes/internal/adapter/http/handlers"
	config "github.com/96vini/n0tes/internal/config"
	postgres "github.com/96vini/n0tes/internal/infrastructure/database/postgres"
	env "github.com/joho/godotenv"
)

func main() {
	// Carrega configuração
	err := env.Load()
	if err != nil {
		log.Fatal("[Database] Erro ao conectar a instancia do banco de dados")
	}

	// Conecta a instancia do banco de dados
	db := connectDatabase(config.DatabaseURL)

	userRepo := postgres.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	router := setupRoutes(userHandler)

	// Inicia servidor
	server.Start(router, cfg.Port)
}
