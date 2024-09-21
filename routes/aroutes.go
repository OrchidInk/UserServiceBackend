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
	user := api.Group("/user", hd.Authorize, hd.SuperAdminOnly)
	user.Post("/register/admin", hd.RegisterAdmin)
	user.Post("/register/user", hd.RegisterUser)

	user.Get("/list/admin", hd.GetListAdmin)
	user.Get("/list/user", hd.GetListUser)

	// category request
	category := api.Group("/categoryEn", hd.Authorize, hd.SuperAdminLogin)
	// create category request
	category.Post("/create/mn", hd.CreateCategoryMn)
	category.Post("/create/en", hd.CreateCategoryEn)

	// update category request
	category.Put("/en/:id", hd.UpdateCategoryEn)
	category.Put("/mn/:id", hd.CreateCategoryMn)

	// delete category request
	category.Delete("/en/:id", hd.DeleteCategoryEn)
	category.Delete("/mn/:id", hd.DeleteCategoryMn)

	// Get category
	category.Get("/list/en", hd.GetListCategoryEn)
	category.Get("/list/mn", hd.GetListCategoryMn)
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

	// Category
	category := api.Group("/category", hd.Authorize)
	category.Get("/list/mn", hd.GetListCategoryMn)
	category.Get("/list/en", hd.GetListCategoryEn)
}
