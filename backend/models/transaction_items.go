package models

import "time"

type TransactionItem struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	TransactionID uint `json:"transaction_id"`
	ProductID     uint `json:"product_id"`

	Quantity  int `json:"quantity"`
	UnitPrice int `json:"unit_price"`
	Subtotal  int `json:"subtotal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Transaction *Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
	Product     *Product     `json:"product" gorm:"foreignKey:ProductID"`
}