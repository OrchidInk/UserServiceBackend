package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
	"orchid.admin.service/utils"
)

func (hd *Handlers) RegisterSuperAdmin(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var superAdmin models.SuperAdminRegister
	if err := ctx.BodyParser(&superAdmin); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := validate.Struct(superAdmin); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "error": err.Error()})
	}

	passwordHash, err := utils.HashPassword(superAdmin.Password)
	if err != nil {
		slog.Error("Unable to hash password", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error"})
	}

	createSuperAdmin, err := queries.CreateUser(ctx.Context(), db.CreateUserParams{
		LastName:         superAdmin.LastName,
		FirstName:        superAdmin.FirstName,
		UserName:         superAdmin.Username,
		Email:            superAdmin.Email,
		IsHashedPassword: passwordHash,
		IsAdmin:          false,
		IsUser:           false,
		IsSuperAdmin:     true,
		IsActive:         true,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // Unique violation error code in PostgreSQL
				if pqErr.Constraint == "User_UserName_key" {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Username already exists"})
				}
				if pqErr.Constraint == "User_Email_key" {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email already exists"})
				}
			}
		}
		slog.Error("Unable to create SuperAdmin", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to create user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "SuperAdmin created", "SuperAdminId": createSuperAdmin.ID})
}

func (hd *Handlers) RegisterAdmin(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var admin models.AdminRegister
	if err := ctx.BodyParser(&admin); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	if err := validate.Struct(admin); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "error": err.Error()})
	}

	passwordHash, err := utils.HashPassword(admin.Password)
	if err != nil {
		slog.Error("Unable to hash password", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Password cannot be hashed"})
	}

	createAdmin, err := queries.CreateUser(ctx.Context(), db.CreateUserParams{
		LastName:         admin.LastName,
		FirstName:        admin.FirstName,
		UserName:         admin.Username,
		Email:            admin.Email,
		IsHashedPassword: passwordHash,
		IsAdmin:          true,
		IsUser:           false,
		IsSuperAdmin:     false,
		IsActive:         true,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				if pqErr.Constraint == "User_UserName_key" {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Username already exists"})
				}
				if pqErr.Constraint == "User_Email_key" {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email already exists"})
				}
			}
		}
		slog.Error("Unable to create admin", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to create admin"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Admin created", "admin_id": createAdmin.ID})
}

func (hd *Handlers) RegisterUser(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var user models.UserRegister
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := validate.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "error": err.Error()})
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		slog.Error("Unable to hash password", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error"})
	}

	createUser, err := queries.CreateUser(ctx.Context(), db.CreateUserParams{
		LastName:         user.LastName,
		FirstName:        user.FirstName,
		UserName:         user.Username,
		Email:            user.Email,
		IsHashedPassword: passwordHash,
		IsAdmin:          false,
		IsUser:           true,
		IsSuperAdmin:     false,
		IsActive:         true,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				if pqErr.Constraint == "User_UserName_key" {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Username already exists"})
				}
				if pqErr.Constraint == "User_Email_key" {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email already exists"})
				}
			}
		}
		slog.Error("Unable to create user", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to create user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created", "user_id": createUser.ID})
}
