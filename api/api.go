package api

import (
	"github.com/gofiber/fiber/v2"

	routes "github.com/ntwaliheritier/go_blog_app/routes"
)

func App() *fiber.App {
	
	app := fiber.New()
	routes.PostRoutes(app.Group("/posts"))
	return app
}