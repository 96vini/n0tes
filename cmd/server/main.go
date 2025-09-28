func main() {
	cfg := config.Load()

	db := connectDB(config.DatabaseURL)

	userRepo := postgres.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	router := setupRoutes(userHandler)

	server.Start(router, cfg.Port)
}