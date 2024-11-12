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
		"message":      "Category created successfully",
		"categoryId":   category.CategoryEnID,
		"categoryName": category.CategoryNameEn,
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
		"message":      "Category created successfully",
		"categoryId":   category.CategoryMnID,
		"categoryName": category.CategoryNameMn,
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
		"message":      "Category updated successfully",
		"categoryId":   updatedCategory.CategoryEnID,
		"categoryName": updatedCategory.CategoryNameEn,
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
		"message":      "Category updated successfully",
		"CategoryID":   updatedCategory.CategoryMnID,
		"categoryName": updatedCategory.CategoryNameMn,
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

func (hd *Handlers) GetCategoriesWithSubCategoriesEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	rows, err := queries.GetCategoriesWithSubCategories(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to fetch categories"})
	}

	categoryMap := make(map[int32]models.CategoryEn)

	for _, row := range rows {
		if _, exists := categoryMap[row.CategoryEnID]; !exists {
			categoryMap[row.CategoryEnID] = models.CategoryEn{
				CategoryEnID:   row.CategoryEnID,
				CategoryNameEn: row.CategoryNameEn,
			}
		}

		if row.SubCategoryIDEn.Valid {
			subCategory := models.SubCategoryEn{
				SubCategoryIdEN:   row.SubCategoryIDEn.Int32,
				SubCategoryNameEN: row.SubCategoryNameEn.String,
				CategoryEnID:      row.CategoryEnID,
			}

			categoryMap[row.CategoryEnID] = appendToSubCategoriesEn(categoryMap[row.CategoryEnID], subCategory)
		}
	}

	var categoryList []fiber.Map
	for _, category := range categoryMap {
		categoryList = append(categoryList, fiber.Map{
			"categoryEnId":   category.CategoryEnID,
			"categoryNameEn": category.CategoryNameEn,
			"subcategories":  category.SubCategories,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(categoryList)
}

func (hd *Handlers) GetCategoriesWithSubCategoriesMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	rows, err := queries.GetCategoriesWithSubCategoriesMn(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to fetch categories"})
	}

	categoryMap := make(map[int32]models.CategoryMn)

	for _, row := range rows {
		if _, exists := categoryMap[row.CategoryMnID]; !exists {
			categoryMap[row.CategoryMnID] = models.CategoryMn{
				CategoryMnID:   row.CategoryMnID,
				CategoryNameMn: row.CategoryNameMn,
			}
		}

		if row.SubCategoryIDMn.Valid {
			subCategory := models.SubCategoryMn{
				SubCategoryIdMn:   row.SubCategoryIDMn.Int32,
				SubCategoryNameMn: row.SubCategoryNameMn.String,
				CategoryMnID:      row.CategoryMnID,
			}

			categoryMap[row.CategoryMnID] = appendToSubCategoriesMn(categoryMap[row.CategoryMnID], subCategory)
		}
	}

	var categoryList []fiber.Map
	for _, category := range categoryMap {
		categoryList = append(categoryList, fiber.Map{
			"categoryMnId":   category.CategoryMnID,
			"categoryNameMn": category.CategoryNameMn,
			"subcategories":  category.SubCategories,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(categoryList)
}

func appendToSubCategoriesEn(category models.CategoryEn, subCategory models.SubCategoryEn) models.CategoryEn {
	if category.SubCategories == nil {
		category.SubCategories = []models.SubCategoryEn{}
	}
	category.SubCategories = append(category.SubCategories, subCategory)
	return category
}

func appendToSubCategoriesMn(category models.CategoryMn, subcategory models.SubCategoryMn) models.CategoryMn {
	if category.SubCategories == nil {
		category.SubCategories = []models.SubCategoryMn{}
	}
	category.SubCategories = append(category.SubCategories, subcategory)
	return category
}
