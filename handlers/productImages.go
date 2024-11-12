package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateProductImagesEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateImageEnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	productImagesEn, err := queries.CreateImagesEn(ctx.Context(), db.CreateImagesEnParams{
		ImagePath:   request.ImagePath,
		ProductEnID: int32(request.ProductEnID),
	})
	if err != nil {
		slog.Error("unable to create product Images", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Id": productImagesEn.Id, "message": "successfully created product images"})
}

func (hd *Handlers) CreateProductImagesMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateImageMnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid request"})
	}

	productImagesMn, err := queries.CreateImageMn(ctx.Context(), db.CreateImageMnParams{
		ProductMnID: int32(request.ProductMnID),
		ImagePath:   request.ImagePath,
	})
	if err != nil {
		slog.Error("unable to create product images mn", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"id": productImagesMn.Id, "message": "successfully created"})
}

func (hd *Handlers) DeleteProductImagesEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid ID"})
	}

	err = queries.DeleteImagesEn(ctx.Context(), int32(id))
	if err != nil {
		slog.Error("unable to delete product image (EN)", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to delete"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully deleted"})
}

func (hd *Handlers) DeleteProductImagesMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid ID"})
	}

	err = queries.DeleteImagesMn(ctx.Context(), int32(id))
	if err != nil {
		slog.Error("unable to delete product image (MN)", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to delete"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully deleted"})
}

func (hd *Handlers) UpdateProductImagesEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateImageEnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	productImage, err := queries.UpdateImagesEn(ctx.Context(), db.UpdateImagesEnParams{
		Id:        int32(request.Id),
		ImagePath: request.ImagePath,
	})
	if err != nil {
		slog.Error("unable to update product image (EN)", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to update"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"id": productImage.Id, "message": "successfully updated"})
}

func (hd *Handlers) UpdateProductImagesMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateImageMnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	productImage, err := queries.UpdateImagesMn(ctx.Context(), db.UpdateImagesMnParams{
		Id:        int32(request.Id),
		ImagePath: request.ImagePath,
	})
	if err != nil {
		slog.Error("unable to update product image (MN)", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to update"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"id": productImage.Id, "message": "successfully updated"})
}

func (hd *Handlers) GetListImagesEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	images, err := queries.GetListImagesEn(ctx.Context())
	if err != nil {
		slog.Error("unable to fetch images (EN)", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to retrieve"})
	}

	return ctx.Status(fiber.StatusOK).JSON(images)
}

func (hd *Handlers) GetListImagesMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	images, err := queries.GetListImagesMn(ctx.Context())
	if err != nil {
		slog.Error("unable to fetch images (MN)", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to retrieve"})
	}

	return ctx.Status(fiber.StatusOK).JSON(images)
}
