package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App) {
	Product(app)
	ApiDoc(app)
}

func ApiDoc(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default
}
