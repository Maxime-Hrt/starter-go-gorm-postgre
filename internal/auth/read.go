package auth

import (
	"net/http"
	"starter-go-gorm-postgresql-fiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var user models.User
	if err := db.Select("id", "username", "email", "created_at", "updated_at").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Database error"})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var users []models.User
	if err := db.Select("id", "username", "email", "created_at", "updated_at").Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch users"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
