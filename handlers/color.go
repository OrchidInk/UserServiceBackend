package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
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
