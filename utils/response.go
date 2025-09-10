package utils

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessfulResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func UnsuccessfulResponse(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(fiber.Map{
				"error": err.Error(),
			})
}