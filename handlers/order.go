package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateOrder(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var req models.CreateOrderRequest
	if err := ctx.BodyParser(&req); err != nil {
		slog.Error("Invalid request body", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid request body"})
	}

	// Use current time if CreatedAt is not set.
	createdAt := req.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now().UTC()
	}

	// Marshal the order items into JSON.
	orderItemsJSON, err := json.Marshal(req.OrderItems)
	if err != nil {
		slog.Error("Failed to marshal order items", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to process order items"})
	}

	// Insert the main order row.
	order, err := queries.CreateOrder(ctx.Context(), db.CreateOrderParams{
		CustomerOrderID: sql.NullInt32{Int32: req.CustomerOrderID, Valid: req.CustomerOrderID != 0},
		CompName:        req.CompName,
		UserName:        req.UserName,
		UserId:          req.UserId,
		PhoneNumber:     req.PhoneNumber,
		OrderItems:      orderItemsJSON,
		CreatedAt:       sql.NullTime{Time: createdAt, Valid: true},
	})
	if err != nil {
		slog.Error("Unable to create order", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to create order"})
	}

	return ctx.Status(fiber.StatusOK).JSON(order)
}

func (hd *Handlers) GetOrdersWithDetails(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	list, err := queries.GetOrdersWithDetails(ctx.Context())
	if err != nil {
		slog.Error("Failed to get orders with details", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}
