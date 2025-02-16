package handlers

import (
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) SendCoins(c *gin.Context) {
	var req struct {
		ToUser string `json:"toUser"`
		Amount int    `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID := c.MustGet("userID").(int)

	if err := h.transactionService.SendCoins(c.Request.Context(), userID, req.ToUser, req.Amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "coins sent successfully"})
}
