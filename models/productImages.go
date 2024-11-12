package models

type ImageMn struct {
	Id          int64  `json:"id"`
	ProductMnID int64  `json:"productMnId" validate:"required"`
	ImagePath   string `json:"imagePath" validate:"required"`
}

type ImageEn struct {
	Id          int64  `json:"id"`
	ProductEnID int64  `json:"productEnId" validate:"required"`
	ImagePath   string `json:"imagePath" validate:"required"`
}

type CreateImageMnRequest struct {
	ProductMnID int64  `json:"productMnId" validate:"required"`
	ImagePath   string `json:"imagePath" validate:"required"`
}

type UpdateImageMnRequest struct {
	Id        int64  `json:"id" validate:"required"`
	ImagePath string `json:"imagePath" validate:"required"`
}

type CreateImageEnRequest struct {
	ProductEnID int64  `json:"productEnId" validate:"required"`
	ImagePath   string `json:"imagePath" validate:"required"`
}

type UpdateImageEnRequest struct {
	Id        int64  `json:"id" validate:"required"`
	ImagePath string `json:"imagePath" validate:"required"`
}
