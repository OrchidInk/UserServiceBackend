package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateDetailEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateDetailEnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid request body"})
	}

	detail, err := queries.CreateDetailEn(ctx.Context(), db.CreateDetailEnParams{
		ProductEnID: request.ProductEnID,
		ChoiceName:  request.ChoiceName,
		ChoiceValue: request.ChoiceValue,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "failed to create detail"})
	}

	return ctx.Status(fiber.StatusOK).JSON(detail)
}
