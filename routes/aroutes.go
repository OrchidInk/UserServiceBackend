package routes

import (
	"github.com/gofiber/fiber/v2"
	"orchid.admin.service/handlers"
)

func superAdminRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/superadmin")
	//
	// SuperAdmin registration and login routes
	api.Post("/register", hd.RegisterSuperAdmin, hd.SuperAdminOnly)
	api.Post("/login", hd.SuperAdminLogin)

	// Routes for registering admins and users
	user := api.Group("/user")
	user.Post("/register/admin", hd.RegisterAdmin)
	user.Post("/register/user", hd.RegisterUser)

	user.Get("/list/superadmin", hd.GetListSuperAdmin)
	user.Get("/list/admin", hd.GetListAdmin)
	user.Get("/list/user", hd.GetListUser)

	//Category
	category := api.Group("/category")
	category.Post("/createEn", hd.CreateCategoryEn)
	category.Post("/createMn", hd.CreateCategoryMn)

	category.Get("/listEn", hd.GetCategoriesWithSubCategoriesEn)
	category.Get("/listMn", hd.GetCategoriesWithSubCategoriesMn)

	category.Patch("/updateEn/:id", hd.UpdateCategoryEn)
	category.Patch("/updateMn/:id", hd.UpdateCategoryMn)

	category.Delete("/deleteEn/:id", hd.DeleteCategoryEn)
	category.Delete("/deleteMn/:id", hd.DeleteCategoryMn)

	category.Get("/with-productEn", hd.GetCategoriesWithSubCategoriesAndProductsEn)
	category.Get("/with-productMn", hd.GetCategoriesWithSubCategoriesAndProductsMn)
	category.Get("/with-productEn/:id", hd.FindSubCategoriesAndProductsByCategoryIDEn)
	category.Get("/with-productMn/:id", hd.FindSubCategoriesAndProductsByCategoryIDMn)

	// subCategory
	subCategory := api.Group("/subCategory")
	subCategory.Post("/createEn", hd.CreateSubCategoryEn)
	subCategory.Post("/createMn", hd.CreateSubCategoryMn)
	subCategory.Get("/list/:SubCategoryIDEn", hd.GetProductsBySubCategoryEn)
	subCategory.Get("/listEn", hd.GetSubCategoryEn)
	subCategory.Get("/listMn", hd.GetSubCategoryMn)
	subCategory.Patch("/updateMn/:id", hd.UpdateSubCategoryMn)
	subCategory.Patch("/updateEn/:id", hd.UpdateSubCategoryEn)
	subCategory.Delete("/deleteMn/:id", hd.DeleteSubCategoryMn)
	subCategory.Delete("/deleteEn/:id", hd.DeleteSubCategoryEn)

	// sCategory
	sCategory := api.Group("/sCategory")
	sCategory.Post("/createEn", hd.CreateSCategoryEn)
	sCategory.Post("/createMn", hd.CreateSCategoryMn)
	sCategory.Get("/listEn", hd.GetAllSCategoryEn)
	sCategory.Get("/listMn", hd.GetAllSCategoryMn)
	sCategory.Patch("/updateEn/:id", hd.UpdateSCategoryEn)
	sCategory.Patch("/updateMn/:id", hd.UpdateSCategoryMn)
	sCategory.Delete("/deleteEn/:id", hd.DeleteBySCategoryEn)
	sCategory.Delete("/deleteMn/:id", hd.DeleteBySCategoryMn)

	sCategory.Get("/listEn/:id", hd.GetProductWithSCategoriesEn)
	sCategory.Get("/listMn/:id", hd.GetProductWithSCategoriesMn)

	//Product
	product := api.Group("/product")
	product.Get("/listEn", hd.GetProductEn)
	product.Get("/listMn", hd.GetProductMn)
	product.Post("/createEn", hd.CreateProductEn)
	product.Post("/createMn", hd.CreateProductMn)

	// product.Get("/listdetailen", hd.GetProductWithDetailsEn)
	// product.Get("/listdetailmn", hd.GetProductWithDetailsMn)
	product.Get("/findEn/:id", hd.FindProductId)
	product.Get("/findMn/:id", hd.FindProductIdMn)

	// Product Update
	product.Put("/purchaseEn/:id", hd.DeductProductStockEn)
	product.Put("/purchaseMn/:id", hd.DeductProductStockMn)
	product.Patch("/updateEn/:id", hd.UpdateProductEn)
	product.Patch("/updateMn/:id", hd.UpdateProductMn)
	// Product Delete
	product.Delete("/deleteEn/:id", hd.DeleteProductEn)
	product.Delete("/deleteMn/:id", hd.DeleteProductMn)

	// Product With Category and subCategory

	//Banner
	banner := api.Group("/banner")
	banner.Post("/create", hd.CreateBanner)
	banner.Get("/list", hd.GetListBanner)
	banner.Put("/update", hd.UpdateBanner)
	banner.Delete("/delete/:id", hd.DeleteBanner)

	// Customer
	customer := api.Group("/customer")
	customer.Post("/create", hd.CreateCustomer)
	customer.Get("/list", hd.GetListCustomer)
	customer.Put("/update/:id", hd.UpdateCustomer)

	// Delivery
	delivery := api.Group("/delivery")
	delivery.Post("/create", hd.CreateDelivery)
	delivery.Get("/list", hd.GetListDelivery)
	delivery.Put("/update/:id", hd.UpdateDelivery)
	delivery.Delete("/delete/:id", hd.DeleteDelivery)

	// Order

	order := api.Group("/order")
	order.Post("/create", hd.CreateOrderItem)
	order.Get("/list", hd.GetOrderItemsByCustomerOrderID)
	order.Patch("/update/:id", hd.UpdateOrderItem)
	order.Delete("/delete/:orderItemId", hd.DeleteOrderItem)

	// Payments
	payments := api.Group("/payment")
	payments.Patch("/update/:id", hd.UpdatePaymentStatus)
	payments.Get("/list", hd.GetListPayment)
}

func adminRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/admin")

	// Admin login route
	api.Post("/login", hd.AdminLogin)

	// Admin info creation and updates
	adminRegistery := api.Group("/registery", hd.AdminOnly, hd.Authorize)
	adminRegistery.Post("/info", hd.CreateAdminInfo)
	adminRegistery.Put("/info/:user_id", hd.UpdateAdminInfo)
	user := api.Group("/user", hd.Authorize, hd.AdminOnly)
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
	subCategory.Get("/list/:SubCategoryIDEn", hd.GetProductsBySubCategoryEn)
	subCategory.Patch("/updateMn/:id", hd.UpdateSubCategoryMn)
	subCategory.Patch("/updateEn/:id", hd.UpdateSubCategoryEn)

	//Product
	product := api.Group("/product")
	product.Post("/createEn", hd.CreateProductEn)
	product.Post("/createMn", hd.CreateProductMn)

	// Product Update
	product.Put("/purchaseEn/:id", hd.DeductProductStockEn)
	product.Put("/purchaseMn/:id", hd.DeductProductStockMn)

	// Product Delete
	product.Delete("/deleteEn/:id", hd.DeleteCategoryEn)
	product.Delete("/deleteMn/:id", hd.DeleteProductMn)

	//Banner
	banner := api.Group("/banner")
	banner.Post("/create", hd.CreateBanner)

	delivery := api.Group("/delivery")
	delivery.Post("/create", hd.CreateDelivery)
	delivery.Get("/list", hd.GetListDelivery)
	delivery.Put("/update/:id", hd.UpdateDelivery)
	delivery.Delete("/delete/:id", hd.DeleteDelivery)

	order := api.Group("/order")
	order.Post("/create", hd.CreateOrderItem)
	order.Get("/list", hd.GetOrderItemsByCustomerOrderID)
	order.Put("/update", hd.UpdateOrderItem)
	order.Delete("/delete/:orderItemId", hd.DeleteOrderItem)
}

func userRoutes(app *fiber.App, hd *handlers.Handlers) {
	api := app.Group("/api/v1/user")

	// User registration and login routes
	api.Post("/register", hd.RegisterUser)
	api.Post("/login", hd.UserLogin)

	// User info creation and updates
	userRegistery := api.Group("/registery")
	userRegistery.Post("/info", hd.CreateUserInfo)
	userRegistery.Put("/info/:user_id", hd.UpdateUserInfo)

	// Category
	category := api.Group("/category")
	category.Get("/listEn", hd.GetCategoriesWithSubCategoriesEn)
	category.Get("/listMn", hd.GetCategoriesWithSubCategoriesMn)

	// Product
	product := api.Group("/product")
	product.Put("/purchaseEn/:id", hd.DeductProductStockEn)
	product.Put("/purchaseMn/:id", hd.DeductProductStockMn)

	delivery := api.Group("/delivery")
	delivery.Post("/create", hd.CreateDelivery)
	delivery.Get("/list", hd.GetListDelivery)
	delivery.Put("/update/:id", hd.UpdateDelivery)
	delivery.Delete("/delete/:id", hd.DeleteDelivery)

	order := api.Group("/order")
	order.Post("/create", hd.CreateOrderItem)
	order.Get("/list", hd.GetOrderItemsByCustomerOrderID)
	order.Put("/update", hd.UpdateOrderItem)
	order.Delete("/delete/:orderItemId", hd.DeleteOrderItem)

	// Payments
	payments := api.Group("/payment")
	payments.Post("/create", hd.CreatePayment)
}
