package controllers

import (
	"backend/helpers"
	"backend/services"
	"backend/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StockIn(c *gin.Context) {
	var req structs.StockInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	Stock, err := services.StockIn(req)

	




}