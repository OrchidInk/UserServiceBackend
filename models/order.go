package models

import (
	"time"
)

type OrderItem struct {
	ProductMnID   int32  `json:"productMnId"`
	ProductEnID   int32  `json:"productEnId"`
	ProductName   string `json:"productName"`
	Quantity      int32  `json:"quantity"`
	PriceAtOrder  string `json:"priceAtOrder"`
	SelectedColor string `json:"selectedColor"`
	SelectedSize  string `json:"selectedSize"`
}

type CreateOrderRequest struct {
	CustomerOrderID int32       `json:"customerOrderId"`
	UserId          int32       `json:"userId"`
	PhoneNumber     string      `json:"phoneNumber"`
	OrderItems      []OrderItem `json:"orderItems"`
	CreatedAt       time.Time   `json:"CreatedAt"`
}
type CreateOrderItemRequest struct {
	OrderID       int32  `json:"orderId"`
	ProductMnID   int32  `json:"productMnId"`
	ProductEnID   int32  `json:"productEnId"`
	Quantity      int32  `json:"quantity"`
	PriceAtOrder  string `json:"priceAtOrder"`
	SelectedColor string `json:"selectedColor"`
	SelectedSize  string `json:"selectedSize"`
}
