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

func FindProduct(c *gin.Context) {
	products, err := services.FindProduct()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get products",
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Products",
		Data:    products,
	})
}

func FindProductById(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product Found",
		Data: structs.ProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
			CategoryId: product.CategoryId,
			SupplierId: product.SupplierId,
			Description: product.Description,
			Total: product.Total,
		},
	})
}

func CreateProduct(c *gin.Context) {
	// Struct request
	var req structs.ProductCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Panggil service
	product, err := services.CreateProduct(req)

	if err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product cannot created",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product updated successfully",
		Data: structs.ProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
			CategoryId: product.CategoryId,
			SupplierId: product.SupplierId,
			Description: product.Description,
			Total: product.Total,
		},
	})

}

func UpdateProduct(c *gin.Context) {

	// Ambil ID product dari parameter URL
	id := c.Param("id")

	// Struct request
	var req structs.ProductUpdateRequest

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
	product, err := services.UpdateProduct(id, req)

	if err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product updated successfully",
		Data: structs.ProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
			CategoryId: product.CategoryId,
			SupplierId: product.SupplierId,
			Description: product.Description,
			Total: product.Total,
		},
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteProduct(id); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to delete product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product deleted successfully",
	})
}
