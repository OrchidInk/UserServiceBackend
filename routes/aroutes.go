package routes

import (
	"github.com/gofiber/fiber/v2"
	"orchid.admin.service/handlers"
)

func superAdminRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/superadmin")
	api.Post("/register", hd.RegisterSuperAdmin, hd.SuperAdminOnly)
	api.Post("/login", hd.SuperAdminLogin, hd.AdminOnly)

	user := api.Group("/register", hd.Authorize, hd.SuperAdminOnly)
	user.Post("/admin", hd.RegisterAdmin, hd.SuperAdminOnly, hd.Authorize)
	user.Post("/user", hd.RegisterUser, hd.SuperAdminOnly, hd.Authorize)
}

func adminRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/admin")
	api.Post("/login", hd.AdminLogin)
}

func userRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/user")

	api.Post("/register", hd.RegisterUser)
	api.Post("/login", hd.UserLogin)
}
