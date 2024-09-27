package handlers

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/shopspring/decimal" // import decimal package
// 	db "orchid.admin.service/db/sqlc"
// 	"orchid.admin.service/models"
// )

// func (hd *Handlers) CreateProductEn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	var req models.CreateProductEnRequest
// 	if err := ctx.BodyParser(&req); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
// 	}

// 	// Parse PriceEn as decimal from string
// 	price, err := decimal.NewFromString(req.PriceEn)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid price format"})
// 	}

// 	product, err := queries.CreateProductEn(ctx.Context(), db.CreateProductEnParams{
// 		CategoryEnID:  req.CategoryEnID,
// 		PriceEn:       price.String(),
// 		StockQuantity: req.StockQuantity,
// 		ImagesPathEn:  req.ImagesPathEn,
// 	})

// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create product"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(product)
// }

// func (hd *Handlers) UpdateProductEnPrice(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	var req models.UpdateProductEnPriceRequest
// 	if err := ctx.BodyParser(&req); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
// 	}

// 	// Parse PriceEn as decimal from string
// 	price, err := decimal.NewFromString(req.PriceEn)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid price format"})
// 	}

// 	product, err := queries.UpdateByProductEnPrice(ctx.Context(), db.UpdateByProductEnPriceParams{
// 		PriceEn:     price.String(),
// 		ProductEnID: req.ProductEnID,
// 	})

// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update product price"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(product)
// }
