package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) SetAdminPermissions(ctx *fiber.Ctx) error {
	isSuperAdmin := ctx.Locals("isSuperAdmin").(bool)

	if !isSuperAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "only super admin can assign permission",
		})
	}

	var permissionRequest models.PermissionRequest
	if err := ctx.BodyParser(&permissionRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := validate.Struct(permissionRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "error": err.Error()})
	}

	queries, _, _ := hd.queries()
	err := queries.SetPermissions(ctx.Context(), db.SetPermissionsParams{
		AdminID:   permissionRequest.AdminID,
		CanCreate: permissionRequest.CanCreate,
		CanRead:   permissionRequest.CanRead,
		CanUpdate: permissionRequest.CanUpdate,
		CanDelete: permissionRequest.CanDelete,
	})
	if err != nil {
		slog.Error("Failed to set permissions", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to set permissions"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Permissions set successfully"})
}

func (hd *Handlers) CheckAdminPermission(ctx *fiber.Ctx, action string) error {
	adminID := ctx.Locals("adminId").(int32)

	queries, _, _ := hd.queries()
	permissions, err := queries.GetPermissionsByAdminID(ctx.Context(), adminID)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Permissions not found"})
	}

	switch action {
	case "create":
		if !permissions.CanCreate.Valid || !permissions.CanCreate.Bool {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Not allowed to create resources"})
		}
	case "read":
		if !permissions.CanRead.Valid || !permissions.CanRead.Bool {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Not allowed to read resources"})
		}
	case "update":
		if !permissions.CanUpdate.Valid || !permissions.CanUpdate.Bool {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Not allowed to update resources"})
		}
	case "delete":
		if !permissions.CanDelete.Valid || !permissions.CanDelete.Bool {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Not allowed to delete resources"})
		}
	}

	return ctx.Next()
}
