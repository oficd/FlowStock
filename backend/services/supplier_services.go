package services

import (
	"backend/database"
	"backend/models"
	"backend/structs"
)

func FindSupplier() ([]models.Supplier, error) {
	var Suppliers []models.Supplier

	result := database.DB.Find(&Suppliers)

	if result.Error != nil {
		return nil, result.Error
	}

	return Suppliers, nil
}

func CreateSupplier(req structs.SupplierCreateRequest) (models.Supplier, error) {
	Supplier := models.Supplier{
		Name: req.Name,
	}

	if err := database.DB.Create(&Supplier).Error; err != nil {
		return Supplier, err
	}

	return Supplier, nil
}

func UpdateSupplier(id string, req structs.SupplierUpdateRequest) (models.Supplier, error) {
	var Supplier models.Supplier

	// Cari Supplier berdasarkan ID
	if err := database.DB.First(&Supplier, id).Error; err != nil {
		return Supplier, err
	}

	// Update data Supplier
	Supplier.Name = req.Name

	// Simpan perubahan
	if err := database.DB.Save(&Supplier).Error; err != nil {
		return Supplier, err
	}

	return Supplier, nil
}

func DeleteSupplier(id string) error {
	if err := database.DB.Delete(&models.Supplier{}, id).Error; err != nil {
		return err
	}

	return nil
}