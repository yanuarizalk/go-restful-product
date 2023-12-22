package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.yanuarizal.net/go-restful-product/config"
)

const (
	ERR_ID_QUERY = 100

	ERR_MSG_UNKNOWN = "unknown error"
	ERR_MSG_PAYLOAD = "invalid payload"
)

func success(c *fiber.Ctx, data any, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    data,
		"message": message,
	})
}

func internalError(c *fiber.Ctx, err error, message string) error {
	body := fiber.Map{
		// "error_id": errId,
		"message": message,
	}

	if config.App.IsDevelopment() {
		body["error"] = err
	}

	return c.Status(fiber.StatusInternalServerError).JSON(body)
}

func invalidRequest(c *fiber.Ctx, err error, message string) error {
	body := fiber.Map{
		// "error_id": errId,
		"message": message,
	}

	if config.App.IsDevelopment() {
		body["error"] = err
	}

	return c.Status(fiber.StatusBadRequest).JSON(body)
}
