package router

import (
	"github.com/gofiber/fiber/v2"
	"github.yanuarizal.net/go-restful-product/handler"
)

func Product(app *fiber.App) {
	r := app.Group("/products")
	r.Post("/", handler.CreateProduct)
	r.Get("/", handler.GetProducts)
	r.Get("/:id", handler.GetProduct)
	r.Put("/:id", handler.UpdateProduct)
	r.Delete("/:id", handler.DeleteProduct)
}
