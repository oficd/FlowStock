package services

import (
	"backend/database"
	"backend/models"
	"backend/structs"
)

func FindCategory() ([]models.Category, error) {
	var Categories []models.Category

	result := database.DB.Find(&Categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return Categories, nil
}

func CreateCategory(req structs.CategoryCreateRequest) (models.Category, error) {
	Category := models.Category{
		Name: req.Name,
	}

	if err := database.DB.Create(&Category).Error; err != nil {
		return Category, err
	}

	return Category, nil
}

func UpdateCategory(id string, req structs.CategoryUpdateRequest) (models.Category, error) {
	var Category models.Category

	// Cari Category berdasarkan ID
	if err := database.DB.First(&Category, id).Error; err != nil {
		return Category, err
	}

	// Update data Category
	Category.Name = req.Name

	// Simpan perubahan
	if err := database.DB.Save(&Category).Error; err != nil {
		return Category, err
	}

	return Category, nil
}

func DeleteCategory(id string) error {
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}