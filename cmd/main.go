package main

import (
	"log"
	"meeting_app/configs"
	"meeting_app/configs/postgres"
	"meeting_app/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	port := configs.EnvGetURI("APP_PORT")

	db, err := postgres.NewPostgresDB()
	if err != nil {
		panic("Failed to connect to the database")
	}
	err = postgres.MigrateDatabase(db.DB)
	if err != nil {
		log.Println("Failed to migrate the database")
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "Metting App API",
		AppName:      "Metting App API v0.0.1",
		BodyLimit:    1024 * 64, // 64 KB
	})

	app.Use(logger.New(logger.Config{
		DisableColors: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(map[string]string{"hello": "world"})
	})
	app.Get("/metrics", monitor.New())

	api := app.Group("/api")

	internal.Setup(api, db.DB)

	app.Listen(":" + port)
}
