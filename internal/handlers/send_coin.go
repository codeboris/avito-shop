package handlers

import (
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendCoinHandler struct {
	transactionService *service.TransactionService
}

func NewSendCoinHandler(transactionService *service.TransactionService) *SendCoinHandler {
	return &SendCoinHandler{transactionService: transactionService}
}

func (h *SendCoinHandler) SendCoin(c *gin.Context) {
	var req struct {
		ToUser string `json:"toUser"`
		Amount int    `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	fromUserID := c.GetInt("userID")
	toUserID := 1 // TODO: Получить ID пользователя по имени

	if err := h.transactionService.SendCoins(fromUserID, toUserID, req.Amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "coins sent successfully"})
}
