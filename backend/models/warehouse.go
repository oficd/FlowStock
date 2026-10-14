package models

import "time"

type Warehouse struct {
	Id          int
	Name        string
	Code        string
	Address     string
	Descriction string
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}