package models

type CreateSize struct {
	SizeName string `json:"SizeName"`
}

type UpdateSize struct {
	SizeId   int32  `json:"sizeId"`
	SizeName string `json:"Size"`
}
