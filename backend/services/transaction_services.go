package services

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/structs"
)

func StockIn(req structs.StockInRequest) (*models.Transaction, error) {

	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	refNumber, err := helpers.GenerateReferenceNumber("STOCK_IN")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 1. Buat transaction
	transaction := models.Transaction{
		ReferenceNumber: refNumber,
		TransactionType: "STOCK_IN",
		SupplierID:      &req.SupplierID,
		ToWarehouseID:   &req.ToWarehouseID,
		TransactionDate: req.TransactionDate,
		Notes:           req.Notes,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 2. Buat transaction items
	for _, item := range req.Items {

		subtotal := item.Quantity * item.UnitPrice

		transactionItem := models.TransactionItem{
			TransactionID: transaction.ID,
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
			UnitPrice:     item.UnitPrice,
			Subtotal:      subtotal,
		}

		if err := tx.Create(&transactionItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 3. Commit
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func StockOut(req structs.StockOutRequest) (*models.Transaction, error) {

	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	referenceNumber, err := helpers.GenerateReferenceNumber("SI")
	if err != nil {
		return nil, err
	}

	// 1. Buat transaction
	transaction := models.Transaction{
		ReferenceNumber: referenceNumber,
		TransactionType: "STOCK_OUT",
		FromWarehouseID:   &req.FromWarehouseID,
		TransactionDate: req.TransactionDate,
		Notes:           req.Notes,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 2. Buat transaction items
	for _, item := range req.Items {

		subtotal := item.Quantity * item.UnitPrice
		quantity := -item.Quantity

		transactionItem := models.TransactionItem{
			TransactionID: transaction.ID,
			ProductID:     item.ProductID,
			Quantity:      quantity,
			UnitPrice:     item.UnitPrice,
			Subtotal:      subtotal,
		}

		if err := tx.Create(&transactionItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 3. Commit
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func FindTransaction() {

}

func Adjustment() {

}

func Transfer() {

}
