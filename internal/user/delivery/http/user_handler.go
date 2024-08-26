package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wahidfebr/go-users-api/internal/user/model"
	"github.com/wahidfebr/go-users-api/internal/user/usecase"
	"net/http"
)

const basePath = "/users"

func RegisterUserRoutes(app *fiber.App, uc usecase.UserUseCase) {
	userGroup := app.Group(basePath)

	userGroup.Post("", func(c *fiber.Ctx) error {
		var user model.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		user.ID = uuid.New().String()
		createdUser, err := uc.CreateUser(user)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
		}
		return c.Status(http.StatusCreated).JSON(createdUser)
	})

	userGroup.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		user, err := uc.GetUserByID(id)
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.JSON(user)
	})

	userGroup.Get("", func(c *fiber.Ctx) error {
		users, err := uc.GetAllUsers()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch users"})
		}
		return c.JSON(users)
	})
}
