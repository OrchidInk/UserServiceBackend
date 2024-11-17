package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateProductEnRequest
	if err := ctx.BodyParser(&request); err != nil {
		slog.Error("unable to parse request body", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid request body"})
	}

	price, err := decimal.NewFromString(request.PriceEn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid price format"})
	}

	createProduct, err := queries.CreateProductEn(ctx.Context(), db.CreateProductEnParams{
		ProductNameEn:   request.ProductNameEn,
		SubCategoryIDEn: request.SubCategoryEnID,
		PriceEn:         price.String(),
		StockQuantity:   request.StockQuantity,
		ImagesPathEn:    request.ImagesPathEn,
	})
	if err != nil {
		slog.Error("unable to create product en", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "product created successfully",
		"product ID": createProduct.ProductEnID,
	})
}

func (hd *Handlers) CreateProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateProductMnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	createProduct, err := queries.CreateProductMn(ctx.Context(), db.CreateProductMnParams{
		ProductNameMn:   request.ProductNameMn,
		SubCategoryIDMn: request.SubCategoryMnID,
		PriceMn:         request.PriceMn,
		StockQuantity:   request.StockQuantity,
		ImagesPathMn:    request.ImagesPathMn,
	})
	if err != nil {
		slog.Error("unable to create product request", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "product create successfully", "product Id": createProduct.ProductMnID})
}

func (hd *Handlers) DeleteProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIDStr := ctx.Params("id")
	ProductID, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindByProductEnID(ctx.Context(), int32(ProductID))
	if err != nil {
		slog.Error("unable to find product", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid find product id"})
	}

	err = queries.DeleteByProductEnId(ctx.Context(), int32(ProductID))
	if err != nil {
		slog.Error("unable to delete productEn id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully deleted product ID"})
}

func (hd *Handlers) DeleteProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIDStr := ctx.Params("id")
	ProductId, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindByProductIdMn(ctx.Context(), int32(ProductId))
	if err != nil {
		slog.Error("unable to find product ID", slog.Any("err", err))
	}

	err = queries.DeleteByProductMnId(ctx.Context(), int32(ProductId))
	if err != nil {
		slog.Error("unable to find product ID", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully deleted product "})
}

func (hd *Handlers) DeductProductStockEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIDStr := ctx.Params("id")
	ProductID, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid product ID"})
	}

	var request struct {
		QuantityPurchased int32 `json:"quantityPurchased"`
	}
	if err := ctx.BodyParser(&request); err != nil || request.QuantityPurchased <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid quantity"})
	}

	updatedProduct, err := queries.DeductStockQuantityByProductEnID(ctx.Context(), db.DeductStockQuantityByProductEnIDParams{
		ProductEnID:       int32(ProductID),
		QuantityPurchased: request.QuantityPurchased,
	})
	if err != nil {
		slog.Error("unable to deduct stock", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to deduct stock", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Stock deducted successfully",
		"productId": updatedProduct.ProductEnID,
		"stockLeft": updatedProduct.StockQuantity,
	})
}

func (hd *Handlers) DeductProductStockMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIDStr := ctx.Params("id")
	ProductID, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		slog.Error("unable to parse id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	var request struct {
		QuantityPurchased int32 `json:"quantityPurchased"`
	}
	if err := ctx.BodyParser(&request); err != nil || request.QuantityPurchased <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid quantity"})
	}

	updatedProduct, err := queries.DeductSockQuantityByProductMnID(ctx.Context(), db.DeductSockQuantityByProductMnIDParams{
		ProductMnID:       int32(ProductID),
		QuantityPurchased: request.QuantityPurchased,
	})
	if err != nil {
		slog.Error("unable to deduct stock", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to deduct stock", "error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Stock deducted successfully",
		"productId": updatedProduct.ProductMnID,
		"stockLeft": updatedProduct.StockQuantity,
	})
}

func (hd *Handlers) GetProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	Product, err := queries.GetListProductEn(ctx.Context())
	if err != nil {
		slog.Error("unable to product get list")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(Product)
}

func (hd *Handlers) GetProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	Product, err := queries.GetListProductMn(ctx.Context())
	if err != nil {
		slog.Error("unable to product get list")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(Product)
}
