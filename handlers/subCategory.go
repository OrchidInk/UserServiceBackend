package handlers

import (
	"database/sql"
	"log/slog"
	"strconv"

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

	// Check if the subcategory already exists by name, not ID
	_, err := queries.FindByNameSubCategoryEn(ctx.Context(), request.SubCategoryNameEN) // Correct column name here
	if err == nil {
		// Subcategory already exists, return conflict
		slog.Error("this subcategory already exists", slog.Any("err", err))
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"err": "subcategory already exists"})
	} else if err != sql.ErrNoRows {
		// Some other error occurred
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	// Proceed with subcategory creation
	createSubCategory, err := queries.CreateSubCategoryEn(ctx.Context(), db.CreateSubCategoryEnParams{
		SubCategoryNameEn: request.SubCategoryNameEN, // Correct column name here
		CategoryEnID:      request.CategoryEnID,
	})
	if err != nil {
		slog.Error("unable to create subCategory", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":           "successfully created",
		"subCategoryID":     createSubCategory.SubCategoryIDEn,
		"subcategoryNameEn": createSubCategory.SubCategoryNameEn,
	})
}

func (hd *Handlers) CreateSubCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.SubCategoryMn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err := queries.FindByNameSubCategoryMn(ctx.Context(), request.SubCategoryNameMn)
	if err == nil {
		slog.Error("this subcategory already exists", slog.Any("err", err))
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"err": "subcategory already exists"})
	} else if err != sql.ErrNoRows {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	createdSubCategory, err := queries.CreateSubCategoryMn(ctx.Context(), db.CreateSubCategoryMnParams{
		SubCategoryNameMn: request.SubCategoryNameMn,
		CategoryMnID:      request.CategoryMnID,
	})
	if err != nil {
		slog.Error("unable to create subcategory", slog.Any("err", err)) // Fix this line
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":         "successfully created",
		"subCategoryId":   createdSubCategory.SubCategoryIDMn,
		"subCategoryName": createdSubCategory.SubCategoryNameMn,
	})
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

func (hd *Handlers) DeleteSubCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	SubCategoryIDStr := ctx.Params("id")
	SubCategoryId, err := strconv.Atoi(SubCategoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid category ID"})
	}

	_, err = queries.FindBySubCategoryIDEn(ctx.Context(), int32(SubCategoryId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}

	err = queries.DeleteSubCategoryEn(ctx.Context(), int32(SubCategoryId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully deleted subcategory"})
}

func (hd *Handlers) DeleteSubCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	SubCategoryIDStr := ctx.Params("id")
	SubCategoryID, err := strconv.Atoi(SubCategoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Category not found"})
	}

	_, err = queries.FindBySubCategoryID(ctx.Context(), int32(SubCategoryID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": err})
	}

	err = queries.DeleteBySubCategoryMn(ctx.Context(), int32(SubCategoryID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "invalid delete subcategory"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "deleted successfully"})
}

func (hd *Handlers) GetProductsBySubCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	subCategoryIDStr := ctx.Params("subCategoryIDEn")
	subCategoryID, err := strconv.Atoi(subCategoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid subcategory ID"})
	}

	products, err := queries.GetProductsBySubCategoryEn(ctx.Context(), int32(subCategoryID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch products"})
	}

	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (hd *Handlers) GetSubCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	cat, err := queries.GetListAllSubCategoriesEn(ctx.Context())
	if err != nil {
		slog.Error("unable to fetching", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(cat)
}

func (hd *Handlers) GetSubCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	cat, err := queries.GetListAllSubCategoryMn(ctx.Context())
	if err != nil {
		slog.Error("unable to fetching", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(cat)
}
