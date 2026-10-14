package models

import "time"

type Product struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Total       int       `json:"total"`

	CategoryId uint      `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryId"`

	SupplierId uint      `json:"supplier_id"`
	Supplier   Supplier  `json:"supplier" gorm:"foreignKey:SupplierId"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}