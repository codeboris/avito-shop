package handlers

import (
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userID := c.MustGet("userID").(int)

	user, err := h.userService.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	purchases, err := h.purchaseService.GetPurchasesByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	transactions, err := h.transactionService.GetTransactionsByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	response := gin.H{
		"coins":     user.Coins,
		"inventory": purchases,
		"coinHistory": gin.H{
			"received": transactions.Received,
			"sent":     transactions.Sent,
		},
	}

	c.JSON(http.StatusOK, response)
}
