package routes

import "github.com/gofiber/fiber/v2"

func PostRoutes(router fiber.Router) {
	router.Get("/", func (ctx *fiber.Ctx) error {
		return ctx.SendString("Hello")
	})
}