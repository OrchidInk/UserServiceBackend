package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
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

func (hd *Handlers) UpdateSize(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	sizeIdStr := ctx.Params("id")
	sizeId, err := strconv.Atoi(sizeIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid category ID"})
	}

	var rqst models.UpdateSize
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err, "ctx": "Bodyparser"})
	}

	_, err = queries.FindByIdSize(ctx.Context(), int32(sizeId))
	if err != nil {
		slog.Error("Cannot find size id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	updatedSize, err := queries.UpdateSize(ctx.Context(), db.UpdateSizeParams{
		SizeId: int32(sizeId),
		Size:   rqst.SizeName,
	})
	if err != nil {
		slog.Error("unable to update this size", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"successfully": "done",
		"Size":         updatedSize.Size,
	})
}

func (hd *Handlers) DeleteSize(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	SizeIdStr := ctx.Params("id")
	SizeId, err := strconv.Atoi(SizeIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid category ID"})
	}

	_, err = queries.FindByIdSize(ctx.Context(), int32(SizeId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": err})
	}

	err = queries.DeleteSize(ctx.Context(), int32(SizeId))
	if err != nil {
		slog.Error("unable to delete size", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "size deleted successfully", "sizeID": SizeId})
}

func (hd *Handlers) GetAllSize(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	size, err := queries.GetAllSize(ctx.Context())
	if err != nil {
		slog.Error("unable to fetch size", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(size)
}

func (hd *Handlers) FindSizeId(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	sizeIDSTR := ctx.Params("id")
	sizeId, err := strconv.Atoi(sizeIDSTR)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err,
		})
	}

	sizeIds, err := queries.FindByIdSize(ctx.Context(), int32(sizeId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(sizeIds)
}
