package models

import "time"

type User struct {
	ID           int32     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	IsAdmin      bool      `json:"is_admin"`
	IsUser       bool      `json:"is_user"`
	IsSuperAdmin bool      `json:"is_super_admin"`
	CreatedAt    time.Time `json:"created_at"`
}

type SuperAdminRegister struct {
	LastName  string `json:"lastname" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type AdminRegister struct {
	LastName  string `json:"lastname" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type UserRegister struct {
	LastName  string `json:"lastname" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type SuperAdminLogin struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminLogin struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
