package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func (hd *Handlers) ListProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	images, err := queries.GetAllProductImagesEn(ctx.Context())
	if err != nil {
		slog.Error("unable to fetching product images", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(images)
}

func (hd *Handlers) ListProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	images, err := queries.GetProductImagesMn(ctx.Context())
	if err != nil {
		slog.Error("unable to fetching product Mn", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(images)
}
