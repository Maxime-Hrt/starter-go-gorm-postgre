package auth

import (
	"net/http"
	"starter-go-gorm-postgresql-fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UpdateUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	validate := c.Locals("validator").(*validator.Validate)

	id := c.Params("id")
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Database error"})
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse request body"})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		if err := user.SetPassword(req.Password); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
		}
	}

	if err := db.Save(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}
