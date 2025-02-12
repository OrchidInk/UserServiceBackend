package models

type Payments struct {
	PaymentID     int32  `json:"paymentID"`
	OrderId       int32  `json:"orderID"`
	UserId        int32  `json:"userID"`
	PaymentMethod string `json:"paymentMethod"`
	PaymentStatus string `json:"paymentStatus"`
	Amount        string `json:"amount"`
}

type CreatePaymentRequest struct {
	OrderId       int32  `json:"orderID"`
	UserId        int32  `json:"userID"`
	PaymentMethod string `json:"paymentMethod"`
	PaymentStatus string `json:"paymentStatus"`
	Amount        string `json:"amount"`
}

type UpdatePaymentStatus struct {
	PaymentStatus string `json:"paymentStatus"`
}
