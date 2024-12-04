package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateDelivery(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateDeliveryRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	createDelivery, err := queries.CreateDelivery(ctx.Context(), db.CreateDeliveryParams{
		DeliverName: request.DeliverName,
		OrderId:     request.OrderId,
	})
	if err != nil {
		slog.Error("unable to request", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "CreatedSuccessfully", "id": createDelivery.DeliverId})

}

func (hd *Handlers) UpdateDelivery(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	updateIDSTR := ctx.Params("Id")
	updateId, err := strconv.Atoi(updateIDSTR)
	if err != nil {
		slog.Error("unablet to parse delivery id", slog.Any("err", "delivery id cannot parse"))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	var request models.UpdateDeliveryRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	delivery, err := queries.FindByDeliveryId(ctx.Context(), int32(updateId))
	if err != nil {
		slog.Error("unable to find delivery", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	updatedDeliver, err := queries.UpdateDeliver(ctx.Context(), db.UpdateDeliverParams{
		DeliverName: request.DeliverName,
		DeliverId:   delivery.DeliverId,
	})
	if err != nil {
		slog.Error("unable to update delivery", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "succesfully", "Id": updatedDeliver.DeliverId})
}

func (hd *Handlers) GetListDelivery(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	list, err := queries.GetListDeliver(ctx.Context())
	if err != nil {
		slog.Error("unable to get list delivery", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

func (hd *Handlers) DeleteDelivery(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	DeliveryIDSTR := ctx.Params("id")

	DeliveryId, err := strconv.Atoi(DeliveryIDSTR)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "unable to find delivery id"})
	}

	delivery, err := queries.FindByDeliveryId(ctx.Context(), int32(DeliveryId))
	if err != nil {
		slog.Error("unable to find delivery id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	err = queries.DeleteDelivery(ctx.Context(), int32(DeliveryId))
	if err != nil {
		slog.Error("unable to deleted delivery", slog.Any("Err", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "successfully deleted", "id": delivery.DeliverId})
}
