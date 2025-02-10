package handlers

import (
	"database/sql"
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateSCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var rqst models.SCategoryEn
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid rqst body"})
	}

	_, err := queries.FindByNameSCategoryNameEn(ctx.Context(), rqst.SCategoryNameEn)
	if err != nil {
		slog.Error("this sCategory already exists", slog.Any("err", err))
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"err": err})
	} else if err != sql.ErrNoRows {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	createSCategory, err := queries.CreateSCategoryEn(ctx.Context(), db.CreateSCategoryEnParams{
		SCategoryNameEn: rqst.SCategoryNameEn,
		SubCategoryIDEn: rqst.SubCategoryIDEn,
	})
	if err != nil {
		slog.Error("unable to create scategory", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":         "successfully created",
		"sCategoryId":     createSCategory.SCategoryIdEn,
		"sCategoryNameEn": createSCategory.SCategoryNameEn,
	})
}

func (hd *Handlers) CreateSCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var rqst models.SCategoryMn
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid rqst body"})
	}

	_, err := queries.FindBySCategoryNameMn(ctx.Context(), rqst.SCategoryName)
	if err != nil {
		slog.Error("this sCategory already created", slog.Any("Err", err))
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"err": err})
	} else if err != sql.ErrNoRows {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	createSCategory, err := queries.CreateSCategoryMn(ctx.Context(), db.CreateSCategoryMnParams{
		SCategoryName:   rqst.SCategoryName,
		SubCategoryIDMn: rqst.SCategoryIdMn,
	})
	if err != nil {
		slog.Error("unable to create sCategory", slog.Any("err", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":         "successfully created",
		"sCategoryIdMn":   createSCategory.SCategoryIdMn,
		"sCategoryNameMn": createSCategory.SCategoryName,
	})
}

func (hd *Handlers) UpdateSCategoryEn(c *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	sCategoryIdStr := c.Params("id")
	sCategoryId, err := strconv.Atoi(sCategoryIdStr)
	if err != nil {
		slog.Error("unable to convert sCategory id", slog.Any("err", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	var rqst models.UpdateSCategoryEn
	if err := c.BodyParser(&rqst); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindBySCategoryIdEn(c.Context(), int32(sCategoryId))
	if err != nil {
		slog.Error("unable to sCategory Id", slog.Any("err", err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": "sCategory not found"})
	}

	updtSCategory, err := queries.UpdateSCategoryNameEn(c.Context(), db.UpdateSCategoryNameEnParams{
		SCategoryNameEn: rqst.SCategoryNameEn,
		SCategoryIdEn:   int32(sCategoryId),
	})
	if err != nil {
		slog.Error("unable to create SCategory", slog.Any("err", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":           "successfully",
		"sCategoryName": updtSCategory.SCategoryNameEn,
		"sCategoryId":   updtSCategory.SubCategoryIDEn,
	})
}

func (hd *Handlers) UpdateSCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	sCategoryIdStr := ctx.Params("id")
	sCategoryId, err := strconv.Atoi(sCategoryIdStr)
	if err != nil {
		slog.Error("unable to convert sCategoryId", slog.Any("Err", err))
	}

	var rqst models.UpdateSCategoryMn
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindBySCategoryIdMn(ctx.Context(), int32(sCategoryId))
	if err != nil {
		slog.Error("not found category id", slog.Any("err", err))
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": err})
	}

	updtSCategory, err := queries.UpdateBySubCategoryNameMn(ctx.Context(), db.UpdateBySubCategoryNameMnParams{
		SubCategoryNameMn: rqst.SCategoryName,
		SubCategoryIDMn:   int32(sCategoryId),
	})
	if err != nil {
		slog.Error("unable to updt sCategory", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":           "successfully created",
		"sCategoryName": updtSCategory.SubCategoryNameMn,
		"sCategoryId":   updtSCategory.CategoryMnID,
	})
}

func (hd *Handlers) DeleteBySCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	sCategoryIdStr := ctx.Params("id")
	sCategoryId, err := strconv.Atoi(sCategoryIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Cant parse sCategoryId"})
	}

	_, err = queries.FindBySCategoryIdEn(ctx.Context(), int32(sCategoryId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": "sCategory not found"})
	}

	err = queries.DeleteSCategoryEn(ctx.Context(), int32(sCategoryId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "cannot delete"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": sCategoryId})
}

func (hd *Handlers) DeleteBySCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	sCategoryMnIdStr := ctx.Params("id")
	sCategoryMnId, err := strconv.Atoi(sCategoryMnIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"err": "cannot parsing id"})
	}

	_, err = queries.FindBySCategoryIdMn(ctx.Context(), int32(sCategoryMnId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": "Cannot find scategory id"})
	}

	err = queries.DeleteSCategoryMn(ctx.Context(), int32(sCategoryMnId))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg id": sCategoryMnId})
}

func (hd *Handlers) GetAllSCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	cat, err := queries.GetAllSCategoryEn(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(cat)
}

func (hd *Handlers) GetAllSCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	cat, err := queries.GetAllSCategoriesMn(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(cat)
}
