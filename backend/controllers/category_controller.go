package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/services"
	"backend/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCategory(c *gin.Context) {
	Categories, err := services.FindCategory()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get Categories",
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Categories",
		Data:    Categories,
	})
}

func FindCategoryById(c *gin.Context) {
	id := c.Param("id")

	var Category models.Category

	if err := database.DB.First(&Category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Category not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Category Found",
		Data: structs.CategoryResponse{
			Id:        Category.Id,
			Name:      Category.Name,
			CreatedAt: Category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: Category.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func CreateCategory(c *gin.Context) {
	// Struct request
	var req structs.CategoryCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Panggil service
	Category, err := services.CreateCategory(req)

	if err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Category cannot created",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Category updated successfully",
		Data: structs.CategoryResponse{
			Id:        Category.Id,
			Name:      Category.Name,
			CreatedAt: Category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: Category.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})

}

func UpdateCategory(c *gin.Context) {

	// Ambil ID Category dari parameter URL
	id := c.Param("id")

	// Struct request
	var req structs.CategoryUpdateRequest

	// Bind JSON request ke struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Panggil service
	Category, err := services.UpdateCategory(id, req)

	if err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Category not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Category updated successfully",
		Data: structs.CategoryResponse{
			Id:        Category.Id,
			Name:      Category.Name,
			CreatedAt: Category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: Category.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteCategory(id); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to delete Category",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Category deleted successfully",
	})
}