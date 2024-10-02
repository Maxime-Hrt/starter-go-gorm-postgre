package api

import (
	"starter-go-gorm-postgresql-fiber/internal/auth"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthController(app *fiber.App, validate *validator.Validate) {
	authGroup := app.Group("/users")

	authGroup.Post("/", func(c *fiber.Ctx) error {
		return auth.CreateUser(c, validate)
	})
}
