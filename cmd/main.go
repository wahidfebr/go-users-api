package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wahidfebr/go-users-api/internal/config"
	"github.com/wahidfebr/go-users-api/internal/user/delivery/http"
	"github.com/wahidfebr/go-users-api/internal/user/repository"
	"github.com/wahidfebr/go-users-api/internal/user/usecase"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New()

	userRepo := repository.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepo)
	http.RegisterUserRoutes(app, userUseCase)

	address := cfg.AppHost + ":" + cfg.AppPort
	log.Printf("Server listening on port http://%s", address)
	err := app.Listen(address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
