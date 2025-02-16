package handlers

import (
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BuyHandler struct {
	merchService *service.MerchService
}

func NewBuyHandler(merchService *service.MerchService) *BuyHandler {
	return &BuyHandler{merchService: merchService}
}

func (h *BuyHandler) BuyItem(c *gin.Context) {
	item := c.Param("item")
	userID := c.GetInt("userID")

	merchID := 1 // TODO: Получить ID товара по имени

	if err := h.merchService.BuyItem(userID, merchID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item purchased successfully"})
}
