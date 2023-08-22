package internal

import (
	internal "meeting_app/internal/app"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app fiber.Router, db *gorm.DB) {
	internal.Run(app, db)
}
