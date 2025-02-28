package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateSize(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var rqst models.CreateSize
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	size, err := queries.CreateSize(ctx.Context(), rqst.SizeName)
	if err != nil {
		slog.Error("unable to create size", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(size)
}
