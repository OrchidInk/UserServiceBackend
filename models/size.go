package models

type CreateSize struct {
	SizeName string `json:"sizeName"`
}

type UpdateSize struct {
	SizeId   int32  `json:"sizeId"`
	SizeName string `json:"sizeName"`
}
