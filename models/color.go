package models

type CreateColor struct {
	ColorName string `json:"ColorName"`
}

type ColorUpdate struct {
	ColorId   int32  `json:"Id"`
	ColorName string `json:"Color"`
}
