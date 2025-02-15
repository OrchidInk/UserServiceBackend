package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateOrderItem(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateOrderItemRequest
	if err := ctx.BodyParser(&request); err != nil {
		slog.Error("Invalid request body", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid request body"})
	}

	price, err := decimal.NewFromString(request.PriceAtOrder)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}
	if price.Exponent() < -2 || price.GreaterThan(decimal.NewFromInt(9999999999)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Price exceeds allowed range (maximum: 9999999999.99)",
		})
	}

	orderItem, err := queries.CreateOrderItem(ctx.Context(), db.CreateOrderItemParams{
		CustomerOrderID: request.CustomerOrderID,
		ProductMnID:     request.ProductMnID,
		ProductEnID:     request.ProductEnID,
		UserId:          request.UserId,
		PhoneNumber:     request.PhoneNumber,
		Quantity:        request.Quantity,
		PriceAtOrder:    price.String(),
	})
	if err != nil {
		slog.Error("Unable to create order item", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to create order item"})
	}

	return ctx.Status(fiber.StatusOK).JSON(orderItem)
}

func (hd *Handlers) GetOrderItemsByCustomerOrderID(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	customerOrderIDStr := ctx.Params("customerOrderId")
	customerOrderID, err := strconv.Atoi(customerOrderIDStr)
	if err != nil {
		slog.Error("Invalid customer order ID", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid customer order ID"})
	}

	orderItems, err := queries.GetOrderItemsByCustomerOrderID(ctx.Context(), int32(customerOrderID))
	if err != nil {
		slog.Error("Unable to fetch order items", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to fetch order items"})
	}

	return ctx.Status(fiber.StatusOK).JSON(orderItems)
}

func (hd *Handlers) UpdateOrderItem(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	OrderIdStr := ctx.Params("id")
	OrderId, err := strconv.Atoi(OrderIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	var request models.UpdateOrderItemRequest
	if err := ctx.BodyParser(&request); err != nil {
		slog.Error("Invalid request body", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid request body"})
	}

	_, err = queries.FindByOrderItemsId(ctx.Context(), int32(OrderId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": err})
	}

	updatedOrderItem, err := queries.UpdateOrderItem(ctx.Context(), db.UpdateOrderItemParams{
		OrderItemID:  int32(OrderId),
		Quantity:     request.Quantity,
		PriceAtOrder: request.PriceAtOrder,
	})
	if err != nil {
		slog.Error("Unable to update order item", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to update order item"})
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedOrderItem)
}

func (hd *Handlers) DeleteOrderItem(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	orderItemIDStr := ctx.Params("orderItemId")
	orderItemID, err := strconv.Atoi(orderItemIDStr)
	if err != nil {
		slog.Error("Invalid order item ID", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid order item ID"})
	}

	err = queries.DeleteOrderItem(ctx.Context(), int32(orderItemID))
	if err != nil {
		slog.Error("Unable to delete order item", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to delete order item"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Order item deleted successfully"})
}
