package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wahidfebr/go-users-api/internal/config"
	"github.com/wahidfebr/go-users-api/internal/container"
	"github.com/wahidfebr/go-users-api/internal/user/delivery/http"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New()

	container := container.NewContainer()

	http.RegisterUserRoutes(app, container.UserUseCase)

	address := cfg.AppHost + ":" + cfg.AppPort
	log.Printf("Server listening on port http://%s", address)
	err := app.Listen(address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
