package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aryamaahendra/mesiakad/pkgs/api"
	"github.com/aryamaahendra/mesiakad/pkgs/api/handlers"
	"github.com/aryamaahendra/mesiakad/pkgs/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	api.NewAPI(app, database.New()).Init()

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))))
}
