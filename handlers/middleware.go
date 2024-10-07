package handlers

import (
	"errors"
	"strings"
	"time"

	"orchid.admin.service/utils"
	"orchid.admin.service/utils/secure"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func (hd *Handlers) Authorize(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": utils.ErrInvalidAuth,
		})
	}

	token := parts[1]

	claims, isAdmin, isSuperAdmin, err := secure.VerifyToken(token, hd.kp)
	if err != nil {
		slog.Error("unable to verify token", slog.Any("err", err))
		return utils.CustomAuthErrMsg(utils.ErrInvalidAuth)
	}

	adminId, ok := claims["adminId"].(float64)
	if !ok {
		return utils.CustomAuthErrMsg("Invalid token claims: adminId")
	}

	expiresAt, ok := claims["exp"].(float64)
	if !ok {
		return utils.CustomAuthErrMsg("Invalid token claims: exp")
	}

	if int64(expiresAt) <= time.Now().UTC().Unix() {
		slog.Error("token expired", slog.Any("err", err))
		return utils.CustomAuthErrMsg(utils.ErrExpiredAuth)
	}

	c.Locals("adminId", int32(adminId))
	c.Locals("isAdmin", isAdmin)
	c.Locals("isSuperAdmin", isSuperAdmin)

	return c.Next()
}

func (hd *Handlers) AuthorizeGetAdminId(c *fiber.Ctx) (int32, error) {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return 0, errors.New("no header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("invalid header format")
	}

	token := parts[1]

	claims, _, _, err := secure.VerifyToken(token, hd.kp)
	if err != nil {
		slog.Error("unable to verify token", slog.Any("err", err))
		return 0, errors.New("invalid token")
	}

	adminId, ok := claims["adminId"].(float64)
	if !ok {
		slog.Error("token claims are missing or incorrect")
		return 0, errors.New("invalid token claims: adminId")
	}

	return int32(adminId), nil
}

func (hd *Handlers) AdminOnly(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Authorization header format"})
	}

	token := parts[1]
	_, isAdmin, _, err := secure.VerifyToken(token, hd.kp)
	if err != nil {
		slog.Error("Invalid token", slog.Any("err", err))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	if !isAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden: Admin access only"})
	}

	return c.Next()
}

func (hd *Handlers) SuperAdminOnly(c *fiber.Ctx) error {
	// Get the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: Missing token",
		})
	}

	// Split the Bearer token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: Invalid token format",
		})
	}

	// Get the token from header
	token := parts[1]

	// Verify the token and check if the user is a SuperAdmin
	claims, _, isSuperAdmin, err := secure.VerifyToken(token, hd.kp)
	if err != nil || !isSuperAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden: SuperAdmin access only",
		})
	}

	// Store relevant claims in the request context (optional, if needed later)
	c.Locals("superAdminId", claims["adminId"])

	// Continue to the next handler
	return c.Next()
}
