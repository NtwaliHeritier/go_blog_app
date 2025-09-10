package routes

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/ntwaliheritier/go_blog_app/models"
	"gorm.io/gorm"
)

func PostRoutes(router fiber.Router, db *gorm.DB) {
	router.Get("/", func (ctx *fiber.Ctx) error {
		var posts []models.Post
		if err := db.Find(&posts).Error; err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(posts)
	})

	router.Post("/", func (ctx *fiber.Ctx) error {
		post := new(models.Post)
		if err := ctx.BodyParser(&post); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := db.Save(post).Error; err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(post)
	})
}