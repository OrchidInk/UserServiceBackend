package handlers

import (
	"log/slog"
	"strconv"

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

func (hd *Handlers) CreateDetailMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateDetailMnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	detail, err := queries.CreateDetailMn(ctx.Context(), db.CreateDetailMnParams{
		ProductMnId: request.ProductMnID,
		ChoiceName:  request.ChoiceName,
		ChoiceValue: request.ChoiceValue,
	})
	if err != nil {
		slog.Error("unable to create detail mn", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(detail)
}

func (hd *Handlers) UpdateDetailEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateDetailEnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid request"})
	}

	detail, err := queries.FindByDetailEn(ctx.Context(), request.DetailEnID)
	if err != nil {
		slog.Error("unable to find detail id", slog.Any("err", err))
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"err": err})
	}

	updatedDetail, err := queries.UpdateDetailMn(ctx.Context(), db.UpdateDetailMnParams{
		ChoiceName:  request.ChoiceName,
		ChoiceValue: request.ChoiceValue,
		DetailMnId:  detail.DetailEnId,
	})
	if err != nil {
		slog.Error("unable to update detail", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "updated succesfuulyy", "detailEnID": updatedDetail.DetailMnId})
}

func (hd *Handlers) UpdateDetailMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateDetailMnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalied request"})
	}

	detail, err := queries.FindByDetailMnID(ctx.Context(), request.DetailMnID)
	if err != nil {
		slog.Error("unable to find detail mn id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	updatedDetail, err := queries.UpdateDetailMn(ctx.Context(), db.UpdateDetailMnParams{
		ChoiceName:  request.ChoiceName,
		ChoiceValue: request.ChoiceValue,
		DetailMnId:  detail.DetailMnId,
	})
	if err != nil {
		slog.Error("unable to updated", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "updated successfully", "detail mn id": updatedDetail.DetailMnId})
}

func (hd *Handlers) DeleteDetailEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	detailIDStr := ctx.Params("id")
	detailID, err := strconv.Atoi(detailIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid this request"})
	}

	_, err = queries.FindByDetailEn(ctx.Context(), int32(detailID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": err})
	}

	err = queries.DeleteDetailEn(ctx.Context(), int32(detailID))
	if err != nil {
		slog.Error("unable to delete detail id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Detail deleted successfully", "detailIdEN": detailID})
}

func (hd *Handlers) DeleteDetailMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	detailIDStr := ctx.Params("id")
	detailID, err := strconv.Atoi(detailIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindByDetailMnID(ctx.Context(), int32(detailID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	err = queries.DeleteDetailMn(ctx.Context(), int32(detailID))
	if err != nil {
		slog.Error("unable to update detail mn", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "deleted successfully", "detail mn id": detailID})
}
