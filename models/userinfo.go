package models

import "time"

type UserInfo struct {
	UserInfoId    int32     `json:"user_info_id"`
	UserId        int32     `json:"user_id"`
	UserImagePath string    `json:"user_image_path"`
	LastName      string    `json:"last_name"`
	FirstName     string    `json:"first_name"`
	Email         string    `json:"email"`
	BirthDate     time.Time `json:"birth_date"`
	PhoneNumber1  string    `json:"phone_number_1"`
	PhoneNumber2  string    `json:"phone_number_2"`
	Address1      string    `json:"address_1"`
}

type UpdateUserInfoRequest struct {
	LastName     string `json:"last_name" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	BirthDate    string `json:"birth_date" validate:"required"`
	PhoneNumber1 string `json:"phone_number_1" validate:"required"`
	PhoneNumber2 string `json:"phone_number_2" validate:"omitempty"`
	Address1     string `json:"address_1" validate:"required"`
}
