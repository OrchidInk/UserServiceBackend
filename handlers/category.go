package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CategoryEn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	_, err := queries.FindByNameCategoryEn(ctx.Context(), request.CategoryNameEn)
	if err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Category already exists"})
	}

	category, err := queries.CreateCategoryEn(ctx.Context(), request.CategoryNameEn)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Category created successfully",
		"category_id":   category.CategoryEnID,
		"category_name": category.CategoryNameEn,
	})
}

func (hd *Handlers) CreateCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CategoryMn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	_, err := queries.FindByNameMnCategoryMn(ctx.Context(), request.CategoryNameMn)
	if err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "category already exist"})
	}

	category, err := queries.CreateCategoryMn(ctx.Context(), request.CategoryNameMn)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Category created successfully",
		"category_id":   category.CategoryMnID,
		"category_name": category.CategoryNameMn,
	})
}

func (hd *Handlers) UpdateCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	var request models.UpdateCategoryEn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	category, err := queries.FindByCategoryEnId(ctx.Context(), request.CategoryEnId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}

	updatedCategory, err := queries.UpdateCategoryEn(ctx.Context(), db.UpdateCategoryEnParams{
		CategoryNameEn: request.CategoryNameEn,
		CategoryEnID:   category.CategoryEnID,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update category"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Category updated successfully",
		"category_id":   updatedCategory.CategoryEnID,
		"category_name": updatedCategory.CategoryNameEn,
	})
}

func (hd *Handlers) UpdateCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateCategoryMn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	category, err := queries.FindByCategoryMnId(ctx.Context(), request.CategoryMnId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "category not found"})
	}

	updatedCategory, err := queries.UpdateCategoryMn(ctx.Context(), db.UpdateCategoryMnParams{
		CategoryNameMn: request.CategoryNameMn,
		CategoryMnID:   category.CategoryMnID,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to updated"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Category updated successfully",
		"Category_ID":   updatedCategory.CategoryMnID,
		"category_name": updatedCategory.CategoryNameMn,
	})
}

func (hd *Handlers) DeleteCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	categoryIDStr := ctx.Params("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid category ID"})
	}

	_, err = queries.FindByCategoryEnId(ctx.Context(), int32(categoryID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}

	err = queries.DeleteCategoryById(ctx.Context(), int32(categoryID))
	if err != nil {
		slog.Error("unable to delete category ID", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Category deleted successfull", "category_ID": categoryID})
}

func (hd *Handlers) DeleteCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	categoryIDStr := ctx.Params("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid category ID"})
	}

	_, err = queries.FindByCategoryMnId(ctx.Context(), int32(categoryID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}

	err = queries.DeleteCategoryByMnId(ctx.Context(), int32(categoryID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete category"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Category deleted successfully", "category_id": categoryID})
}

func (hd *Handlers) GetListCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	category, err := queries.GetListByAllCategoryEn(ctx.Context())
	if err != nil {
		slog.Error("unable to get list", slog.Any("err", "err"))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "err list"})
	}

	return ctx.JSON(category)
}

func (hd *Handlers) GetListCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	category, err := queries.GetListByAllCategoryMn(ctx.Context())
	if err != nil {
		slog.Error("unable to get list", slog.Any("err", "err"))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "err get list"})
	}

	return ctx.JSON(category)
}
