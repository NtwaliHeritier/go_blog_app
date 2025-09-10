package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	routes "github.com/ntwaliheritier/go_blog_app/routes"
)

func App(db *gorm.DB) *fiber.App {
	
	app := fiber.New()
	routes.PostRoutes(app.Group("/posts"), db)
	return app
}