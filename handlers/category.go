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

	var createCategoryEn models.CreateCategoryEn
	if err := ctx.BodyParser(&createCategoryEn); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid this request"})
	}

	if err := validate.Struct(createCategoryEn); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "cannot be sent null"})
	}

	_, err := queries.FindByNameCategoryEn(ctx.Context(), createCategoryEn.CategoryNameEn)
	if err != nil {
		slog.Error("already created category name ", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "already created this category"})
	}

	CreateCategoryEn, err := queries.CreateCategoryEn(ctx.Context(), createCategoryEn.CategoryNameEn)
	if err != nil {
		slog.Error("unable to create category en", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to create"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "categoryEn created", "CategoryId": CreateCategoryEn.CategoryEnID})
}

func (hd *Handlers) CreateCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var createCategoryMn models.CreateCategoryMn
	if err := ctx.BodyParser(&createCategoryMn); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	if err := validate.Struct(createCategoryMn); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "cannot be null data sent"})
	}

	_, err := queries.FindByNameCategoryMn(ctx.Context(), createCategoryMn.CategoryNameMn)
	if err != nil {
		slog.Error("already created this category Mn", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Already created category"})
	}

	CreateCategoryMn, err := queries.CreateCategoryMn(ctx.Context(), createCategoryMn.CategoryNameMn)
	if err != nil {
		slog.Error("unable to create this category", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Err"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "created categoryMn", "categoryMNId": CreateCategoryMn.CategoryMnID})
}

func (hd *Handlers) UpdateCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var updateRequestEn models.UpdateCategoryEn
	if err := ctx.BodyParser(&updateRequestEn); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid this request"})
	}

	if err := validate.Struct(updateRequestEn); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "cannot sent nil body"})
	}

	categoryEN, err := queries.FindByNameCategoryEn(ctx.Context(), updateRequestEn.CategoryNameEn)
	if err != nil {
		slog.Error("unable to find category en", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	updateCategoryEn, err := queries.UpdateCategoryEn(ctx.Context(), db.UpdateCategoryEnParams{
		CategoryNameEn: updateRequestEn.CategoryNameEn,
		CategoryEnID:   categoryEN.CategoryEnID,
	})
	if err != nil {
		slog.Error("unable to update category", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully", "categoryEnId": updateCategoryEn.CategoryEnID})
}

func (hd *Handlers) UpdateCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var updateRequestMn models.UpdateCategoryMn
	if err := ctx.BodyParser(&updateRequestMn); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	if err := validate.Struct(updateRequestMn); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "err"})
	}

	categoryMn, err := queries.FindByNameCategoryMn(ctx.Context(), updateRequestMn.CategoryNameMn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "err"})
	}

	err = queries.UpdateByCategoryMn(ctx.Context(), db.UpdateByCategoryMnParams{
		CategoryNameMn: updateRequestMn.CategoryNameMn,
		CategoryMnID:   categoryMn.CategoryMnID,
	})
	if err != nil {
		slog.Error("unable to update categoryMN", slog.Any("Err", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully", "categoryMnId": categoryMn.CategoryMnID})
}

func (hd *Handlers) DeleteCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	CategoryIDStr := ctx.Params("id")

	categoryID, err := strconv.Atoi(CategoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	err = queries.DeleteByIdCategoryEn(ctx.Context(), int32(categoryID))
	if err != nil {
		slog.Error("unable to delete category en", slog.Any("Err", "Err"))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "err"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"categoryEnId": categoryID, "message": "successfully deleted"})
}

func (hd *Handlers) DeleteCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	CategoryIDStr := ctx.Params("id")

	categoryID, err := strconv.Atoi(CategoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	err = queries.DeleteFromCategoryMn(ctx.Context(), int32(categoryID))
	if err != nil {
		slog.Error("unable to delete category mn id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"categoryMNID": categoryID, "message": "successfully deleted"})
}

func (hd *Handlers) GetListCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	CategoryEn, err := queries.GetListByAllCategoryEn(ctx.Context())
	if err != nil {
		slog.Error("unable to get list", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.JSON(CategoryEn)
}

func (hd *Handlers) GetListCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	CategoryMn, err := queries.GetListAllCategoryMn(ctx.Context())
	if err != nil {
		slog.Error("unable to get list", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "err"})
	}

	return ctx.JSON(CategoryMn)
}

func (hd *Handlers) FindByIdCategoryEn(ctx *fiber.Ctx) error {
	return nil
}

func (hd *Handlers) FindByIdCategoryMn(ctx *fiber.Ctx) error {
	return nil
}
