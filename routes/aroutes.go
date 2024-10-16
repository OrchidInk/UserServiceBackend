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

	//Category
	category := api.Group("/category")
	category.Post("/createEn", hd.CreateCategoryEn)
	category.Post("/createMn", hd.CreateCategoryMn)

	category.Get("/listEn", hd.GetCategoriesWithSubCategoriesEn)
	category.Get("/listMn", hd.GetCategoriesWithSubCategoriesMn)

	// subCategory
	subCategory := api.Group("/subCategory")
	subCategory.Post("/createEn", hd.CreateSubCategoryEn)
	subCategory.Post("/createMn", hd.CreateSubCategoryMn)
	subCategory.Get("/list/:subCategoryIDEn", hd.GetProductsBySubCategoryEn)

	//Product
	product := api.Group("/product")
	product.Post("/createEn", hd.CreateProductEn)
	product.Post("/createMn", hd.CreateProductMn)

	product.Put("/purchaseEn/:id", hd.DeductProductStockEn)
	product.Put("/purchaseMn/:id", hd.DeductProductStockMn)

	product.Delete("/deleteEn/:id", hd.DeleteCategoryEn)
	product.Delete("/deleteMn/:id", hd.DeleteProductMn)

	file := api.Group("/file")
	file.Post("/create", hd.UploadFile)

	banner := api.Group("/banner")
	banner.Post("/create", hd.CreateBanner)
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
	category.Get("/listEn", hd.GetCategoriesWithSubCategoriesEn)
	category.Get("/listMn", hd.GetCategoriesWithSubCategoriesMn)

	product := api.Group("/product", hd.Authorize)
	product.Put("/purchaseEn/:id", hd.DeductProductStockEn)
	product.Put("/purchaseMn/:id", hd.DeductProductStockMn)
}
