package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateBanner(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateBannerRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	createBanner, err := queries.CreateBannerInfo(ctx.Context(), request.BannerImageUrl)
	if err != nil {
		slog.Error("unable to create banner", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "successfully", "bannerID": createBanner.BannerId})
}

func (hd *Handlers) UpdateBanner(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateBannerRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	banner, err := queries.FindByBannerId(ctx.Context(), request.BannerId)
	if err != nil {
		slog.Error("unable to find banner", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	updateBanner, err := queries.UpdateBannerInfo(ctx.Context(), db.UpdateBannerInfoParams{
		BannerId:       banner.BannerId,
		BannerImageUrl: request.BannerImageUrl,
	})
	if err != nil {
		slog.Error("unable to update banner", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "updated banner", "id": updateBanner.BannerId})
}

func (hd *Handlers) DeleteBanner(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	BannerDeleteIDStr := ctx.Params("id")
	BannerId, err := strconv.Atoi(BannerDeleteIDStr)
	if err != nil {
		slog.Error("unable to parse id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindByBannerId(ctx.Context(), int32(BannerId))
	if err != nil {
		slog.Error("unable to find banner id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	err = queries.DeleteBannerInfo(ctx.Context(), int32(BannerId))
	if err != nil {
		slog.Error("unable to delete banner id", slog.Any("err", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "deleted successfully", "banner Id": BannerId})
}

func (hd *Handlers) GetListBanner(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	banner, err := queries.GetAllBanners(ctx.Context())
	if err != nil {
		slog.Error("unable to find banners", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(banner)
}
