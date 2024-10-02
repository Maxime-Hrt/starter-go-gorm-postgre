package api

import (
	"starter-go-gorm-postgresql-fiber/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthController(app *fiber.App) {
	authGroup := app.Group("/users")

	authGroup.Post("", auth.CreateUser)
	authGroup.Get("/:id", auth.GetUser)
	authGroup.Get("", auth.GetUsers)
	authGroup.Patch("/:id", auth.UpdateUser)
	authGroup.Delete("/:id", auth.DeleteUser)
}
