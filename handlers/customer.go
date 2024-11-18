package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateCustomer(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	var request models.CreateCustomerRequest

	// Parse the request body
	if err := ctx.BodyParser(&request); err != nil {
		slog.Error("Failed to parse request body", "error", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate the request data
	if request.CustomerName == "" || request.ContractStartDate.IsZero() || request.ContractEndDate.IsZero() {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Ensure ContractEndDate is after ContractStartDate
	if request.ContractEndDate.Before(request.ContractStartDate) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Contract end date must be after start date"})
	}

	// Create the customer record
	customer, err := queries.CreateCustomer(ctx.Context(), db.CreateCustomerParams{
		CustomerName:      request.CustomerName,
		ContractStartDate: request.ContractStartDate,
		ContractEndDate:   request.ContractEndDate,
		IsActive:          request.IsActive,
	})
	if err != nil {
		slog.Error("Failed to create customer", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create customer"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Customer created successfully",
		"customerId":    customer.CustomerId,
		"customerName":  customer.CustomerName,
		"contractStart": customer.ContractStartDate.Format(time.RFC3339),
		"contractEnd":   customer.ContractEndDate.Format(time.RFC3339),
		"isActive":      customer.IsActive,
	})
}
