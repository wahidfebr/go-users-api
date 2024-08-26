package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wahidfebr/go-users-api/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	address := cfg.AppHost + ":" + cfg.AppPort
	log.Printf("Server listening on port http://%s", address)
	err := app.Listen(address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
