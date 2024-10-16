package models

import "time"

type Image struct {
	Id        int32     `json:"id"`
	ImagePath string    `json:"image_path"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateImageResponse struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	ImagePath string `json:"image_path"`
}
