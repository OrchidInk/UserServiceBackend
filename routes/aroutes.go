package routes

import (
	"github.com/gofiber/fiber/v2"
	"orchid.admin.service/handlers"
)

func superAdminRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/superadmin")

	// SuperAdmin registration and login routes
	api.Post("/register", hd.RegisterSuperAdmin, hd.SuperAdminOnly)
	api.Post("/login", hd.SuperAdminLogin)

	// Routes for registering admins and users
	user := api.Group("/register", hd.Authorize, hd.SuperAdminOnly)
	user.Post("/admin", hd.RegisterAdmin, hd.SuperAdminOnly, hd.Authorize)
	user.Post("/user", hd.RegisterUser, hd.SuperAdminOnly, hd.Authorize)
}

func adminRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/admin")

	// Admin login route
	api.Post("/login", hd.AdminLogin)

	// Admin info creation and updates
	adminRegistery := api.Group("/registery", hd.AdminOnly, hd.Authorize)
	adminRegistery.Post("/info", hd.CreateAdminInfo)
	adminRegistery.Put("/info/:user_id", hd.UpdateAdminInfo)
}

func userRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/user")

	// User registration and login routes
	api.Post("/register", hd.RegisterUser)
	api.Post("/login", hd.UserLogin)

	// User info creation and updates
	userRegistery := api.Group("/registery", hd.Authorize)
	userRegistery.Post("/info", hd.CreateUserInfo)
	userRegistery.Put("/info/:user_id", hd.UpdateUserInfo)
}
