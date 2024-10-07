package models

type CreateDetailEnRequest struct {
	DetailEnID  int32  `json:"detailEnID"`
	ProductEnID int32  `json:"productEnID"`
	ChoiceName  string `json:"choiceName"`
	ChoiceValue string `json:"choiceValue"`
}
