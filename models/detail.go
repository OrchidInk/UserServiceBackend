package models

type DetailEn struct {
	DetailEnID  int32  `json:"detailEnId"`
	ProductEnID int32  `json:"productEnId"`
	ChoiceName  string `json:"choice_name"`
	ChoiceValue string `json:"choice_value"`
}

type CreateDetailEnRequest struct {
	ProductEnID int32  `json:"productEnId" validate:"required"`
	ChoiceName  string `json:"choiceName" validate:"required"`
	ChoiceValue string `json:"choiceValue" validate:"required"`
}

type CreateDetailMnRequest struct {
	ProductMnID int32  `json:"productMnId" validte:"required"`
	ChoiceName  string `json:"choiceName" validate:"required"`
	ChoiceValue string `json:"choiceValue" validate:"required"`
}

type UpdateDetailEnRequest struct {
	DetailEnID  int32  `json:"detailEnId" validate:"required"`
	ChoiceName  string `json:"choiceName" validate:"required"`
	ChoiceValue string `json:"choiceValue" validate:"required"`
}

type DetailEnResponse struct {
	DetailEnID  int32  `json:"detailEnId"`
	ProductEnID int32  `json:"productEnId"`
	ChoiceName  string `json:"choiceName"`
	ChoiceValue string `json:"choiceValue"`
}
