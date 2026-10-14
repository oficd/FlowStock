package services

import (
	"backend/database"
	"backend/models"
	"backend/structs"
)

func FindProduct() ([]models.Product, error) {
	var products []models.Product

	// preload Category and Supplier so related names are included in responses
	result := database.DB.Preload("Category").Preload("Supplier").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func CreateProduct(req structs.ProductCreateRequest) (models.Product, error) {
	product := models.Product{
		Name: req.Name,
		Description: req.Description,
		Total: req.Total,
		CategoryId: req.CategoryId,
		SupplierId: req.SupplierId,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func UpdateProduct(id string, req structs.ProductUpdateRequest) (models.Product, error) {
	var product models.Product

	// Cari product berdasarkan ID
	if err := database.DB.First(&product, id).Error; err != nil {
		return product, err
	}

	// Update data product
	product.Name = req.Name
	product.Description = req.Description
	product.Total = req.Total
	product.CategoryId = req.CategoryId
	product.SupplierId = req.SupplierId

	// Simpan perubahan
	if err := database.DB.Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func DeleteProduct(id string) error {
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}