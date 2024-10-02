package auth

import (
	"net/http"
	"starter-go-gorm-postgresql-fiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeleteUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete user"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}
