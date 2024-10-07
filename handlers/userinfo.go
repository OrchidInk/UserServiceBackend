package handlers

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateAdminInfo(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var adminInfo models.UserInfo
	if err := ctx.BodyParser(&adminInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := validate.Struct(adminInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "error": err.Error()})
	}

	adminID, err := queries.FindByUserID(ctx.Context(), adminInfo.UserId)
	if err != nil {
		slog.Error("unable to create this admin", slog.Any("err", err))
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid create admin"})
	}

	if !adminID.IsAdmin {
		slog.Error("this user role is not admin", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid this user not admin role user"})
	}

	if !adminID.IsActive {
		slog.Error("unable to active admin", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "deactive admin user cannot be reqistery"})
	}

	_, err = queries.CreateUserInfo(ctx.Context(), db.CreateUserInfoParams{
		UserId:        adminID.ID,
		UserImagePath: adminInfo.UserImagePath,
		LastName:      adminID.LastName,
		FirstName:     adminID.FirstName,
		Email:         adminID.Email,
		BirthDate:     adminInfo.BirthDate,
		PhoneNumber1:  adminInfo.PhoneNumber1,
		PhoneNumber2:  adminInfo.PhoneNumber2,
		Address1:      adminInfo.Address1,
	})
	if err != nil {
		slog.Error("unable to create admin info", slog.Any("err", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Admin info registered successfully",
		"admin_id": adminInfo.UserId,
	})
}

func (hd *Handlers) CreateUserInfo(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var userInfo models.UserInfo
	if err := ctx.BodyParser(&userInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	if err := validate.Struct(userInfo); err != nil {
		slog.Error("cannot be sent nil body", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid struct body"})
	}

	userId, err := queries.FindByUserID(ctx.Context(), userInfo.UserId)
	if err != nil {
		slog.Error("unable to create user", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid this user not registered"})
	}
	if !userId.IsUser {
		slog.Error("unable to create userinfo this user is not user role", slog.Any("message", "invalid request is user info created"))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid not user"})
	}

	if !userId.IsActive {
		slog.Error("unable to active user", slog.Any("message", "this user is deactive"))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "deactive user cannot be create user info"})
	}

	_, err = queries.CreateUserInfo(ctx.Context(), db.CreateUserInfoParams{
		UserId:       userId.ID,
		LastName:     userId.LastName,
		FirstName:    userId.FirstName,
		Email:        userId.Email,
		BirthDate:    userInfo.BirthDate,
		PhoneNumber1: userInfo.PhoneNumber1,
		PhoneNumber2: userInfo.PhoneNumber2,
		Address1:     userInfo.Address1,
	})
	if err != nil {
		slog.Error("unable to create user info registery", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "unable to create userInfo registery"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Admin info registered successfully",
		"user_id": userInfo.UserId,
	})
}

func (hd *Handlers) UpdateAdminInfo(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var updateInfo models.UpdateUserInfoRequest
	if err := ctx.BodyParser(&updateInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := validate.Struct(updateInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "error": err.Error()})
	}

	userId, err := ctx.ParamsInt("user_id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user ID"})
	}

	err = queries.UpdateUserInfo(ctx.Context(), db.UpdateUserInfoParams{
		LastName:     updateInfo.LastName,
		FirstName:    updateInfo.FirstName,
		Email:        updateInfo.Email,
		BirthDate:    time.Now(),
		PhoneNumber1: updateInfo.PhoneNumber1,
		PhoneNumber2: updateInfo.PhoneNumber2,
		Address1:     updateInfo.Address1,
		UserId:       int32(userId),
	})
	if err != nil {
		slog.Error("Failed to update admin info", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update admin info"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Admin info updated successfully",
	})
}

func (hd *Handlers) UpdateUserInfo(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var updateInfo models.UpdateUserInfoRequest
	if err := ctx.BodyParser(&updateInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := validate.Struct(updateInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "error": err.Error()})
	}

	userId, err := ctx.ParamsInt("user_id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user ID"})
	}

	err = queries.UpdateUserInfo(ctx.Context(), db.UpdateUserInfoParams{
		LastName:     updateInfo.LastName,
		FirstName:    updateInfo.FirstName,
		Email:        updateInfo.Email,
		BirthDate:    time.Now(), // Use actual birthdate parsing here
		PhoneNumber1: updateInfo.PhoneNumber1,
		PhoneNumber2: updateInfo.PhoneNumber2,
		Address1:     updateInfo.Address1,
		UserId:       int32(userId),
	})
	if err != nil {
		slog.Error("Failed to update user info", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update user info"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User info updated successfully",
	})
}

