package main

import (
	"log"
	"starter-go-gorm-postgresql-fiber/api"
	"starter-go-gorm-postgresql-fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerValidator struct {
	validator *validator.Validate
}

func (cv *CustomerValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	validate := validator.New()

	dsn := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create the table for the User model if it does not exist
	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	api.AuthController(app, validate)

	log.Fatal(app.Listen(":3000"))
}
