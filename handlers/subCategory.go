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
	return nil
}
