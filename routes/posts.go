package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	models "github.com/ntwaliheritier/go_blog_app/models"
	utils "github.com/ntwaliheritier/go_blog_app/utils"
	"gorm.io/gorm"
)

func PostRoutes(router fiber.Router, db *gorm.DB) {
	router.Get("/", func (ctx *fiber.Ctx) error {
		posts := new([]models.Post)
		if err := db.Find(posts).Error; err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		return utils.SuccessfulResponse(ctx, posts)
	})

	router.Post("/", func (ctx *fiber.Ctx) error {
		post := new(models.Post)
		if err := ctx.BodyParser(&post); err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		if err := db.Save(post).Error; err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		return utils.SuccessfulResponse(ctx, post)
	})

	router.Get("/:id", func (ctx *fiber.Ctx) error {
		post := new(models.Post)
		postId, err := ctx.ParamsInt("id")
		
		if err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		db.Where("id = ?", postId).Find(post)

		if post.ID == 0 {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, fmt.Errorf("post with id %d does not exist", postId))
		}

		return utils.SuccessfulResponse(ctx, post)
	})

	router.Put("/:id", func (ctx *fiber.Ctx) error {
		postId, err := ctx.ParamsInt("id")
		
		if err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		post := new(models.Post)

		if err := db.Where("id = ?", postId).First(&post).Error; err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, fmt.Errorf("post with id %d does not exist", postId))
		}

		if err := ctx.BodyParser(post); err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		if err := db.Save(post).Error; err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		return utils.SuccessfulResponse(ctx, post)
	})

	router.Delete("/:id", func (ctx *fiber.Ctx) error {
		postId, err := ctx.ParamsInt("id")
		post := new(models.Post)
		
		if err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		if err := db.Where("id = ?", postId).First(post).Error; err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, fmt.Errorf("post with id %d does not exist", postId))
		}

		if err := db.Delete(post).Error; err != nil {
			return utils.UnsuccessfulResponse(ctx, fiber.StatusBadRequest, err)
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	})
}