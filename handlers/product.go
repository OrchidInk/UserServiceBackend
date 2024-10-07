package handlers

import (
	"log/slog"

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
		PriceEn:         price.String(), // Using the converted float64 price
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
