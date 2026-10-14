package models

import "time"

type Transaction struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	ReferenceNumber string    `json:"reference_number" gorm:"unique;not null"`
	TransactionType string    `json:"transaction_type" gorm:"not null"`

	SupplierID      *uint `json:"supplier_id"`
	FromWarehouseID *uint `json:"from_warehouse_id"`
	ToWarehouseID   *uint `json:"to_warehouse_id"`

	Reason string `json:"reason"`
	Notes  string `json:"notes"`

	TransactionDate time.Time `json:"transaction_date"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Supplier      *Supplier  `json:"supplier" gorm:"foreignKey:SupplierID"`
	FromWarehouse *Warehouse `json:"from_warehouse" gorm:"foreignKey:FromWarehouseID"`
	ToWarehouse   *Warehouse `json:"to_warehouse" gorm:"foreignKey:ToWarehouseID"`

	Items []TransactionItem `json:"items" gorm:"foreignKey:TransactionID"`
}