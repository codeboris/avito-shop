package handlers

import (
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MerchHandler struct {
	merchService    *service.MerchService
	purchaseService *service.PurchaseService
}

func NewMerchHandler(merchService *service.MerchService, purchaseService *service.PurchaseService) *MerchHandler {
	return &MerchHandler{merchService: merchService, purchaseService: purchaseService}
}

func (h *MerchHandler) BuyMerch(c *gin.Context) {
	merchName := c.Param("item")

	userID := c.MustGet("userID").(int)

	if err := h.purchaseService.BuyMerch(c.Request.Context(), userID, merchName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "merch purchased successfully"})
}
