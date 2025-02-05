package controllers

import (
	"fetchrewards-go/models"
	"fetchrewards-go/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getReceiptPoints)
}

func processReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := services.ProcessReceipt(receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func getReceiptPoints(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}
	points, err := services.GetReceiptPoints(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "receipt not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": points})
}
