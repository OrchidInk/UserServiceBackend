package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreatePayment(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var rqst models.CreatePaymentRequest
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid rqst body"})
	}

	payment, err := queries.CreatePayment(ctx.Context(), db.CreatePaymentParams{
		OrderID:       rqst.OrderId,
		UserID:        rqst.UserId,
		PaymentMethod: rqst.PaymentMethod,
		PaymentStatus: rqst.PaymentStatus,
		Amount:        rqst.Amount,
	})
	if err != nil {
		slog.Error("unable to create payments", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Err": "Cannot create payment"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message":   "Successfully",
		"PaymentId": payment.PaymentID,
		"Date":      payment.CreatedAt.Time.UTC(),
	})
}

func (hd *Handlers) GetListPayment(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	payments, err := queries.GetListPayment(ctx.Context())
	if err != nil {
		slog.Error("unable to fetch payments", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Err": "cannot fetch payments list"})
	}

	return ctx.Status(fiber.StatusOK).JSON(payments)
}

func (hd *Handlers) UpdatePaymentStatus(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	PaymentIdStr := ctx.Params("id")
	PaymentId, err := strconv.Atoi(PaymentIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	var rqst models.UpdatePaymentStatus
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid rqst body"})
	}

	_, err = queries.FindByPaymentsId(ctx.Context(), int32(PaymentId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": "not found payment id"})
	}

	updatedPayment, err := queries.UpdatePaymentStatus(ctx.Context(), db.UpdatePaymentStatusParams{
		PaymentStatus: rqst.PaymentStatus,
		PaymentID:     int32(PaymentId),
	})
	if err != nil {
		slog.Error("unable to update payments status", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "cannot update payment status"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":          "succesffully",
		"updatedPaymentId": updatedPayment.PaymentID,
	})
}
