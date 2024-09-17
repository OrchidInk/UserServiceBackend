package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"orchid.admin.service/handlers"
	"orchid.admin.service/utils"
)

func Routes(hd *handlers.Handlers) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.CustomOverErrMsg,
		BodyLimit:    100 * 1024 * 1024,
	})

	app.Use(cors.New(
		cors.Config{
			AllowMethods: "GET,POST,PATCH,DELETE,PUT",
		},
	))

	superAdminRoutes(app, hd)
	adminRoutes(app, hd)
	userRoutes(app, hd)

	return app
}
