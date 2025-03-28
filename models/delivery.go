package models

import "time"

type CreateDeliveryRequest struct {
	DeliveryId    int32     `json:"deliveryId"`
	DeliverName   string    `json:"deliverName"`
	OrderId       int32     `json:"orderId"`
	DeliverAmount string    `json:"deliverAmount"`
	CreatedAt     time.Time `json:"CreatedAt"`
}

type UpdateDeliveryRequest struct {
	DeliveryId  int32  `json:"deliveryId"`
	DeliverName string `json:"deliverName"`
}
