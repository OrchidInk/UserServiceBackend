package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateColor(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var rqst models.CreateColor
	if err := ctx.BodyParser(&rqst); err != nil {
		slog.Error("Failed to parse rqst body", slog.Any("Err", err))
	}

	color, err := queries.CreateColor(ctx.Context(), rqst.ColorName)
	if err != nil {
		slog.Error("unable to create color", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(color)
}

func (hd *Handlers) UpdateColor(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	colorIdStr := ctx.Params("id")
	colorId, err := strconv.Atoi(colorIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	var rqst models.ColorUpdate
	if err := ctx.BodyParser(&rqst); err != nil {
		slog.Error("Failed to parse rqst body", slog.Any("Err", err))
	}

	_, err = queries.FindByColorId(ctx.Context(), int32(colorId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": err})
	}

	updatedColor, err := queries.UpdateColor(ctx.Context(), db.UpdateColorParams{
		ColorId: int32(colorId),
		Color:   rqst.ColorName,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"successfully": "updated color",
		"ColorName":    updatedColor.Color,
	})
}

func (hd *Handlers) DeleteColor(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	colorIdStr := ctx.Params("id")
	colorId, err := strconv.Atoi(colorIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindByColorId(ctx.Context(), int32(colorId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": err})
	}

	err = queries.DeleteColor(ctx.Context(), int32(colorId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Successully deleted": "true",
		"ColorId":             colorId,
	})
}
