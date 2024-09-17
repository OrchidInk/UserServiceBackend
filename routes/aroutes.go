package routes

import (
	"github.com/gofiber/fiber/v2"
	"orchid.admin.service/handlers"
)

func superAdminRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/superadmin")
	api.Post("/register", hd.RegisterSuperAdmin)
	api.Post("/register/admin", hd.RegisterAdmin)
	api.Post("/register/user", hd.RegisterUser)
}

func adminRoutes(app *fiber.App, hd *handlers.Handlers) {}

func userRoutes(app *fiber.App, hd *handlers.Handlers) {}
