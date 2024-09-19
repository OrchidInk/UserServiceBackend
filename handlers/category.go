package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
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

	return nil
}

func (hd *Handlers) UpdateCategoryMn(ctx *fiber.Ctx) error {

	return nil
}

func (hd *Handlers) DeleteCategoryEn(ctx *fiber.Ctx) error {

	return nil
}

func (hd *Handlers) DeleteCategoryMn(ctx *fiber.Ctx) error {

	return nil
}

func (hd *Handlers) GetListCategoryEn(ctx *fiber.Ctx) error {

	return nil
}

func (hd *Handlers) GetListCategoryMn(ctx *fiber.Ctx) error {

	return nil
}

func (hd *Handlers) FindByIdCategoryEn(ctx *fiber.Ctx) error {
	return nil
}

func (hd *Handlers) FindByIdCategoryMn(ctx *fiber.Ctx) error {
	return nil
}
