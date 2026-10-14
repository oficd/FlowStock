package structs

import "time"

type TransactionItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
	UnitPrice int  `json:"unit_price"`
}

type TransactionItemResponse struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
	UnitPrice int  `json:"unit_price"`
}

type StockInRequest struct {
	SupplierID      uint                     `json:"supplier_id" binding:"required"`
	ToWarehouseID   uint                     `json:"to_warehouse_id" binding:"required"`
	TransactionDate time.Time                `json:"transaction_date"`
	Notes           string                   `json:"notes"`
	Items           []TransactionItemRequest `json:"items" binding:"required,min=1"`
}

type StockOutRequest struct {
	FromWarehouseID uint                     `json:"from_warehouse_id" binding:"required"`
	TransactionDate time.Time                `json:"transaction_date"`
	Reason          string                   `json:"reason"`
	Notes           string                   `json:"notes"`
	Items           []TransactionItemRequest `json:"items" binding:"required,min=1"`
}

type TransferRequest struct {
	FromWarehouseID uint                     `json:"from_warehouse_id" binding:"required"`
	ToWarehouseID   uint                     `json:"to_warehouse_id" binding:"required"`
	TransactionDate time.Time                `json:"transaction_date"`
	Notes           string                   `json:"notes"`
	Items           []TransactionItemRequest `json:"items" binding:"required,min=1"`
}

type AdjustmentItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type AdjustmentRequest struct {
	WarehouseID     uint                     `json:"warehouse_id" binding:"required"`
	TransactionDate time.Time                `json:"transaction_date"`
	Reason          string                   `json:"reason" binding:"required"`
	Notes           string                   `json:"notes"`
	Items           []AdjustmentItemRequest  `json:"items" binding:"required,min=1"`
}

type TransactionResponse struct {
	ID              uint                     `json:"id"`
	ReferenceNumber string                   `json:"reference_number"`
	TransactionType string                   `json:"transaction_type"`

	SupplierID      *uint `json:"supplier_id,omitempty"`
	FromWarehouseID *uint `json:"from_warehouse_id,omitempty"`
	ToWarehouseID   *uint `json:"to_warehouse_id,omitempty"`

	Reason          string                   `json:"reason,omitempty"`
	Notes           string                   `json:"notes,omitempty"`
	TransactionDate time.Time                `json:"transaction_date"`

	Items []TransactionItemResponse `json:"items"`

	CreatedAt time.Time `json:"created_at"`
}