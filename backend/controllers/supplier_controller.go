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

func FindSupplier(c *gin.Context) {
	Suppliers, err := services.FindSupplier()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get Suppliers",
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Suppliers",
		Data:    Suppliers,
	})
}

func FindSupplierById(c *gin.Context) {
	id := c.Param("id")

	var Supplier models.Supplier

	if err := database.DB.First(&Supplier, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Supplier not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Supplier Found",
		Data: structs.SupplierResponse{
			Id:        uint(Supplier.Id),
			Name:      Supplier.Name,
			CreatedAt: Supplier.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: Supplier.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func CreateSupplier(c *gin.Context) {
	// Struct request
	var req structs.SupplierCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Panggil service
	Supplier, err := services.CreateSupplier(req)

	if err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Supplier cannot created",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Supplier updated successfully",
		Data: structs.SupplierResponse{
			Id:        uint(Supplier.Id),
			Name:      Supplier.Name,
			CreatedAt: Supplier.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: Supplier.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})

}

func UpdateSupplier(c *gin.Context) {

	// Ambil ID Supplier dari parameter URL
	id := c.Param("id")

	// Struct request
	var req structs.SupplierUpdateRequest

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
	Supplier, err := services.UpdateSupplier(id, req)

	if err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Supplier not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Supplier updated successfully",
		Data: structs.SupplierResponse{
			Id:        uint(Supplier.Id),
			Name:      Supplier.Name,
			CreatedAt: Supplier.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: Supplier.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteSupplier(id); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to delete Supplier",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Supplier deleted successfully",
	})
}