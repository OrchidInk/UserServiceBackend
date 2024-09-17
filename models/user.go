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

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserInfo struct {
	UserId    int32     `json:"userId"`
	LastName  string    `json:"lastName"`
	FirstName string    `json:"firstName"`
	RegNo     string    `json:"regNo"`
	Email     string    `json:"email"`
	Phone1    string    `json:"phone1"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserInfoRegister struct {
	UserId    int32  `json:"userId"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	RegNo     string `json:"regNo"`
	Email     string `json:"email"`
	Phone1    string `json:"phone1"`
	Address   string `json:"address"`
	BirthDate string `json:"birthDate"`
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
