package models

type OrderItem struct {
	OrderItemID     int32  `json:"orderItemId"`
	CustomerOrderID int32  `json:"customerOrderId"`
	ProductMnID     int32  `json:"productMnId"`
	ProductEnID     int32  `json:"productEnId"`
	UserId          int32  `json:"userId"`
	PhoneNumber     string `json:"phoneNumber"`
	Quantity        int32  `json:"quantity"`
	PriceAtOrder    string `json:"priceAtOrder"`
}

type CreateOrderItemRequest struct {
	CustomerOrderID int32  `json:"customerOrderId"`
	ProductMnID     int32  `json:"productMnId"`
	ProductEnID     int32  `json:"productEnId"`
	UserId          int32  `json:"userId"`
	PhoneNumber     string `json:"phoneNumber"`
	Quantity        int32  `json:"quantity" validate:"required"`
	PriceAtOrder    string `json:"priceAtOrder" validate:"required"`
}

type UpdateOrderItemRequest struct {
	OrderItemID  int32  `json:"orderItemId" validate:"required"`
	Quantity     int32  `json:"quantity" validate:"required"`
	PriceAtOrder string `json:"priceAtOrder" validate:"required"`
}

type OrderItemResponse struct {
	OrderItemID     int32  `json:"orderItemId"`
	CustomerOrderID int32  `json:"customerOrderId"`
	ProductMnID     int32  `json:"productMnId"`
	ProductEnID     int32  `json:"productEnId"`
	Quantity        int32  `json:"quantity"`
	PriceAtOrder    string `json:"priceAtOrder"`
}
