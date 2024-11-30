package handlers

import (
	"database/sql"
	"strconv"
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

func (hd *Handlers) GetListCustomer(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	org, err := queries.GetAllCustomers(ctx.Context())
	if err != nil {
		slog.Error("unable to fetched customer", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Unable to customer fetched"})
	}
	return ctx.Status(fiber.StatusOK).JSON(org)
}

func (hd *Handlers) UpdateCustomer(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	// Extract CustomerId from URL parameter
	customerIDStr := ctx.Params("Id")
	customerId, err := strconv.Atoi(customerIDStr)
	if err != nil || customerId <= 0 {
		slog.Error("Invalid or missing Customer ID in URL", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid Customer ID in URL"})
	}

	// Parse request body
	var rqst models.UpdateCustomerIsActiveRequest
	if err := ctx.BodyParser(&rqst); err != nil {
		slog.Error("Invalid request body", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid request body"})
	}

	// Check if the customer exists in the database
	_, err = queries.FindByCustomerId(ctx.Context(), int32(customerId))
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Customer ID not found", slog.Int("CustomerID", customerId))
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": "Customer ID not found"})
		}
		slog.Error("Database error while finding customer", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Internal server error"})
	}

	// Update customer status
	customerUpdate, err := queries.UpdateCustomerIsActive(ctx.Context(), db.UpdateCustomerIsActiveParams{
		CustomerId:      int32(customerId), // Pass the ID from URL param
		IsActive:        rqst.IsActive,     // From the body
		ContractEndDate: rqst.ContractEndDate,
	})
	if err != nil {
		slog.Error("Unable to update customer", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to update customer"})
	}

	// Respond with success
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Customer updated successfully",
		"customerID": customerUpdate.CustomerId,
	})
}
