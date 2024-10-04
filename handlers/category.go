package handlers

import (
	"github.com/gofiber/fiber/v2"
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
	return nil
}

func (hd *Handlers) DeleteCategoryEn(ctx *fiber.Ctx) error {
	return nil
}

func (hd *Handlers) UpdateCategoryMn(ctx *fiber.Ctx) error {
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
