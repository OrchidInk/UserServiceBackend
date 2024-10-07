package models

type DetailEn struct {
	DetailEnID  int32  `json:"detail_en_id"`
	ProductEnID int32  `json:"product_en_id"`
	ChoiceName  string `json:"choice_name"`
	ChoiceValue string `json:"choice_value"`
}

type CreateDetailEnRequest struct {
	ProductEnID int32  `json:"product_en_id" validate:"required"`
	ChoiceName  string `json:"choice_name" validate:"required"`
	ChoiceValue string `json:"choice_value" validate:"required"`
}

type UpdateDetailEnRequest struct {
	DetailEnID  int32  `json:"detail_en_id" validate:"required"`
	ChoiceName  string `json:"choice_name" validate:"required"`
	ChoiceValue string `json:"choice_value" validate:"required"`
}

type DetailEnResponse struct {
	DetailEnID  int32  `json:"detail_en_id"`
	ProductEnID int32  `json:"product_en_id"`
	ChoiceName  string `json:"choice_name"`
	ChoiceValue string `json:"choice_value"`
}
