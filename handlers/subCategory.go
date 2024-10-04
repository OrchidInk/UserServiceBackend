package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateSubCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.SubCategoryEn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid request body"})
	}

	subCategory, err := queries.FindBySubCategoryIDEn(ctx.Context(), request.SubCategoryIdEN)
	if err != nil {
		slog.Error("this sub category already exist", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	createSubCategory, err := queries.CreateSubCategoryEn(ctx.Context(), db.CreateSubCategoryEnParams{
		SubCategoryNameEN: request.SubCategoryNameEN,
		CategoryEnID:      subCategory.CategoryEnID,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfullly created", "subCategoryID": createSubCategory.SubCategoryIDEn, "subcategoryNameEn": createSubCategory.SubCategoryNameEn})
}

func (hd *Handlers) CreateSubCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.SubCategoryMn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	subCategory, err := queries.FindBySubCategoryID(ctx.Context(), request.SubCategoryIdMn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	createdSubCategory, err := queries.CreateSubCategoryMn(ctx.Context(), db.CreateSubCategoryMnParams{
		SubCategoryNameMn: request.SubCategoryNameMn,
		CategoryMnID:      subCategory.CategoryMnID,
	})
	if err != nil {
		slog.Error("unable to sub category", slog.Any("err", "err"))
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully", "subCategoryId": createdSubCategory.SubCategoryIDMn, "subCategoryName": createdSubCategory.SubCategoryNameMn})
}

func (hd *Handlers) UpdateSubCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateSubCategoryEn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid message request"})
	}

	subCategoryID, err := queries.FindBySubCategoryIDEn(ctx.Context(), request.SubCategoryEnId)
	if err != nil {
		slog.Error("unable to find sub category ID", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	updateSubCategory, err := queries.UpdateSubCategoryNameEn(ctx.Context(), db.UpdateSubCategoryNameEnParams{
		SubCategoryNameEn: request.SubCategoryNameEn,
		SubCategoryIDEn:   subCategoryID.SubCategoryIDEn,
	})
	if err != nil {
		slog.Error("unable to update request", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully", "subCategoryName": updateSubCategory.SubCategoryNameEn, "subCategoryID": updateSubCategory.SubCategoryIDEn})
}

func (hd *Handlers) UpdateSubCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateSubCategoryMn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	subCategory, err := queries.FindBySubCategoryID(ctx.Context(), request.SubCategoryMnID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": "unable to find id"})
	}

	updateSubCategory, err := queries.UpdateBySubCategoryNameMn(ctx.Context(), db.UpdateBySubCategoryNameMnParams{
		SubCategoryNameMn: request.SubCategoryNameMn,
		SubCategoryIDMn:   subCategory.SubCategoryIDMn,
	})
	if err != nil {
		slog.Error("unable to update subCategory", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully", "subCategoryName": updateSubCategory.SubCategoryNameMn, "subCategoryId": updateSubCategory.SubCategoryIDMn})
}
