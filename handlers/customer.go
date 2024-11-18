package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateCustomer(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateCustomerRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid request"})
	}

	contractStartDate := int32(request.ContractStartDate.Unix())
	contractEndDate := int32(request.ContractEndDate.Unix())

	customer, err := queries.CreateCustomer(ctx.Context(), db.CreateCustomerParams{
		CustomerName:      request.CustomerName,
		ContractStartDate: contractStartDate,
		ContractEndDate:   contractEndDate,
		IsActive:          request.IsActive,
	})
	if err != nil {
		slog.Error("unable to create customer", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      customer.CustomerId,
		"message": "Customer successfully created",
	})
}
