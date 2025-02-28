package models

type CreateColor struct {
	ColorName string `json:"colorName"`
}

type ColorUpdate struct {
	ColorId   int32  `json:"colorId"`
	ColorName string `json:"colorName"`
}
