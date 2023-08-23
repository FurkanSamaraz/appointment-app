package main

import (
	"log"
	_ "meeting_app/cmd/docs"
	"meeting_app/configs"
	"meeting_app/configs/postgres"
	"meeting_app/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

// @title          Meetin APP API
// @version        1.0
// @BasePath       /
// @schemes        http https
// @Accept         json
// @Produce        json
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

	app.Get("/metrics", monitor.New())

	api := app.Group("/api")
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json", // Sunucu URL'sine göre güncelleyin
		DeepLinking: false,
	}))

	internal.Setup(api, db.DB)

	app.Listen(":" + port)
}
