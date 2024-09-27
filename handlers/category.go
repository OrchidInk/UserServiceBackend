package handlers

// import (
// 	"log/slog"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// 	db "orchid.admin.service/db/sqlc"
// 	"orchid.admin.service/models"
// )

// func (hd *Handlers) CreateCategoryEn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	var createCategoryEn models.CreateCategoryEn
// 	if err := ctx.BodyParser(&createCategoryEn); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
// 	}

// 	if err := validate.Struct(createCategoryEn); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "validation failed"})
// 	}

// 	_, err := queries.FindByNameCategoryEn(ctx.Context(), createCategoryEn.CategoryNameEn)
// 	if err == nil {
// 		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "category already exists"})
// 	}

// 	category, err := queries.CreateCategoryEn(ctx.Context(), db.CreateCategoryEnParams{
// 		CategoryNameEn:   createCategoryEn.CategoryNameEn,
// 		ParentCategoryID: createCategoryEn.ParentCategoryID,
// 	})
// 	if err != nil {
// 		slog.Error("unable to create category en", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to create category"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "category created", "CategoryId": category.CategoryEnID})
// }

// func (hd *Handlers) CreateCategoryMn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	var createCategoryMn models.CreateCategoryMn
// 	if err := ctx.BodyParser(&createCategoryMn); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
// 	}

// 	if err := validate.Struct(createCategoryMn); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "validation failed"})
// 	}

// 	_, err := queries.FindByNameMnCategoryMn(ctx.Context(), createCategoryMn.CategoryNameMn)
// 	if err == nil {
// 		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "category already exists"})
// 	}

// 	category, err := queries.CreateCategoryMn(ctx.Context(), db.CreateCategoryMnParams{
// 		CategoryNameMn:   createCategoryMn.CategoryNameMn,
// 		ParentCategoryID: createCategoryMn.ParentCategoryID,
// 	})
// 	if err != nil {
// 		slog.Error("unable to create category mn", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to create category"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "category created", "CategoryId": category.CategoryMnID})
// }

// func (hd *Handlers) UpdateCategoryEn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	var updateRequestEn models.UpdateCategoryEn
// 	if err := ctx.BodyParser(&updateRequestEn); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
// 	}

// 	if err := validate.Struct(updateRequestEn); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "validation failed"})
// 	}

// 	categoryEn, err := queries.FindByCategoryEnId(ctx.Context(), updateRequestEn.CategoryEnID)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "category not found"})
// 	}

// 	updatedCategory, err := queries.UpdateCategoryEn(ctx.Context(), db.UpdateCategoryEnParams{
// 		CategoryNameEn:   updateRequestEn.CategoryNameEn,
// 		ParentCategoryID: updateRequestEn.ParentCategoryID,
// 		CategoryEnID:     categoryEn.CategoryEnID,
// 	})
// 	if err != nil {
// 		slog.Error("unable to update category en", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to update category"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "category updated", "CategoryId": updatedCategory.CategoryEnID})
// }

// func (hd *Handlers) UpdateCategoryMn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	var updateRequestMn models.UpdateCategoryMn
// 	if err := ctx.BodyParser(&updateRequestMn); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
// 	}

// 	if err := validate.Struct(updateRequestMn); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "validation failed"})
// 	}

// 	categoryMn, err := queries.FindByCategoryMnId(ctx.Context(), updateRequestMn.CategoryMnID)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "category not found"})
// 	}

// 	updatedCategory, err := queries.UpdateCategoryMn(ctx.Context(), db.UpdateCategoryMnParams{
// 		CategoryNameMn:   updateRequestMn.CategoryNameMn,
// 		ParentCategoryID: updateRequestMn.ParentCategoryID,
// 		CategoryMnID:     categoryMn.CategoryMnID,
// 	})
// 	if err != nil {
// 		slog.Error("unable to update category mn", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to update category"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "category updated", "CategoryId": updatedCategory.CategoryMnID})
// }

// func (hd *Handlers) DeleteCategoryEn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()
// 	categoryIDStr := ctx.Params("id")

// 	categoryID, err := strconv.Atoi(categoryIDStr)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid category ID"})
// 	}

// 	err = queries.DeleteCategoryById(ctx.Context(), int32(categoryID))
// 	if err != nil {
// 		slog.Error("unable to delete category en", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to delete category"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"CategoryId": categoryID, "message": "category deleted"})
// }

// func (hd *Handlers) DeleteCategoryMn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()
// 	categoryIDStr := ctx.Params("id")

// 	categoryID, err := strconv.Atoi(categoryIDStr)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid category ID"})
// 	}

// 	err = queries.DeleteCategoryByMnId(ctx.Context(), int32(categoryID))
// 	if err != nil {
// 		slog.Error("unable to delete category mn", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to delete category"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"CategoryId": categoryID, "message": "category deleted"})
// }

// func (hd *Handlers) GetListCategoryEn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	categories, err := queries.GetListByAllCategoryEn(ctx.Context())
// 	if err != nil {
// 		slog.Error("unable to get categories en", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to fetch categories"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(categories)
// }

// func (hd *Handlers) GetListCategoryMn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	categories, err := queries.GetListByAllCategoryMn(ctx.Context())
// 	if err != nil {
// 		slog.Error("unable to get categories mn", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unable to fetch categories"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(categories)
// }
