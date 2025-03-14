package models

import "time"

type Customer struct {
	CustomerId        int32     `json:"customerId"`
	CustomerName      string    `json:"customerName"`
	ContractStartDate time.Time `json:"contractStartDate"`
	ContractEndDate   time.Time `json:"contractEndDate"`
	IsActive          bool      `json:"isActive"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type CreateCustomerRequest struct {
	CustomerName      string    `json:"customerName" validate:"required"`
	ContractStartDate time.Time `json:"contractStartDate" validate:"required"`
	ContractEndDate   time.Time `json:"contractEndDate" validate:"required"`
	IsActive          bool      `json:"isActive"`
}

type UpdateCustomerIsActiveRequest struct {
	CustomerID      int32     `json:"customerId" validate:"required"`
	ContractEndDate time.Time `json:"contractEndDate" validate:"required"`
	IsActive        bool      `json:"isActive" validate:"required"`
}

type UpdateCustomerContractDateRequest struct {
	CustomerID        int32     `json:"customerId" validate:"required"`
	ContractStartDate time.Time `json:"contractStartDate" validate:"required"`
	ContractEndDate   time.Time `json:"contractEndDate" validate:"required"`
}

type CustomerCountResponse struct {
	ActiveCount   int32 `json:"activeCount"`
	InactiveCount int32 `json:"inactiveCount"`
}
