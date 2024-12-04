package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
	"orchid.admin.service/utils"
	"orchid.admin.service/utils/secure"
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

func (hd *Handlers) SuperAdminLogin(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var loginRequest models.SuperAdminLogin
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	superAdminLogin, err := queries.FindByUserName(ctx.Context(), loginRequest.UserName)
	if err != nil || !superAdminLogin.IsSuperAdmin {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	if !utils.CheckPassword(loginRequest.Password, superAdminLogin.IsHashedPassword) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalied super admin password"})
	}

	token, err := secure.IssueToken(superAdminLogin.ID, superAdminLogin.IsAdmin, superAdminLogin.IsSuperAdmin, hd.kp)
	if err != nil {
		slog.Error("Unable to generate token", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not log in"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"role":    "SuperAdmin",
		"user":    superAdminLogin.UserName,
	})
}

func (hd *Handlers) AdminLogin(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var loginRequest models.AdminLogin
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"messgae": "Invalied request in body"})
	}

	adminLogin, err := queries.FindByUserName(ctx.Context(), loginRequest.UserName)
	if err != nil || !adminLogin.IsAdmin {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized admin"})
	}

	if !utils.CheckPassword(loginRequest.Password, adminLogin.IsHashedPassword) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "password is wrong"})
	}

	token, err := secure.IssueToken(adminLogin.ID, adminLogin.IsSuperAdmin, adminLogin.IsAdmin, hd.kp)
	if err != nil {
		slog.Error("unable to generate token err", slog.Any("err", err))
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unable to generate token this is not admin role user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"role":    "Admin",
		"user":    adminLogin.UserName,
	})
}

func (hd *Handlers) UserLogin(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var loginRequest models.UserLogin
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalied request"})
	}

	userLogin, err := queries.FindByUserName(ctx.Context(), loginRequest.UserName)
	if err != nil || !userLogin.IsUser {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "this user is not user login"})
	}

	if !utils.CheckPassword(loginRequest.Password, userLogin.IsHashedPassword) {
		slog.Error("wrong password err", slog.Any("err", err))
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "wrong password"})
	}

	token, err := secure.IssueToken(userLogin.ID, userLogin.IsSuperAdmin, userLogin.IsAdmin, hd.kp)
	if err != nil {
		slog.Error("unable to generate token", slog.Any("err", err))
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unable to generate token"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "user created",
		"userid":   userLogin.ID,
		"token":    token,
		"username": userLogin.UserName,
	})
}

func (hd *Handlers) GetListAdmin(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	admins, err := queries.GetListAdmin(ctx.Context())
	if err != nil {
		slog.Error("unable to get admin users", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return ctx.JSON(admins)
}

func (hd *Handlers) GetListUser(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	users, err := queries.GetListUser(ctx.Context())
	if err != nil {
		slog.Error("unable to get admin users", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return ctx.JSON(users)
}

func (hd *Handlers) GetListSuperAdmin(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	superAdmin, err := queries.FindBySuperAdminAdmin(ctx.Context())
	if err != nil {
		slog.Error("unable to fetch superadmin")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.JSON(superAdmin)
}
