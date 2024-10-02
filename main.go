package main

import (
	"fmt"
	"log"
	"os"
	"starter-go-gorm-postgresql-fiber/api"
	"starter-go-gorm-postgresql-fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		},
	})

	app.Use(logger.New())

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	var dsn string
	if dbPassword != "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris",
			dbHost, dbUser, dbPassword, dbName, dbPort)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris",
			dbHost, dbUser, dbName, dbPort)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		c.Locals("validator", validator.New())
		return c.Next()
	})

	api.AuthController(app)

	log.Fatal(app.Listen(":3000"))
}
