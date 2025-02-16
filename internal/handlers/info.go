package handler

import (
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoHandler struct {
	userService        *service.UserService
	merchService       *service.MerchService
	transactionService *service.TransactionService
}

func NewInfoHandler(userService *service.UserService, merchService *service.MerchService, transactionService *service.TransactionService) *InfoHandler {
	return &InfoHandler{
		userService:        userService,
		merchService:       merchService,
		transactionService: transactionService,
	}
}

func (h *InfoHandler) GetInfo(c *gin.Context) {
	userID := c.GetInt("userID")

	user, err := h.userService.GetUserInfo(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user info"})
		return
	}

	// Получаем историю транзакций
	transactions, err := h.transactionService.GetTransactionHistory(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get transaction history"})
		return
	}

	// Получаем инвентарь пользователя
	purchases, err := h.merchService.GetPurchasesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get inventory"})
		return
	}

	// Формируем историю монет
	coinHistory := models.CoinHistory{
		Received: []models.CoinTransaction{},
		Sent:     []models.CoinTransaction{},
	}

	for _, t := range transactions {
		if t.ToUser == userID {
			coinHistory.Received = append(coinHistory.Received, models.CoinTransaction{
				User:   "User", // TODO: Получить имя пользователя
				Amount: t.Amount,
			})
		} else if t.FromUser == userID {
			coinHistory.Sent = append(coinHistory.Sent, models.CoinTransaction{
				User:   "User", // TODO: Получить имя пользователя
				Amount: t.Amount,
			})
		}
	}

	// Формируем инвентарь
	var inventory []models.InventoryItem
	for _, p := range purchases {
		merch, err := h.merchService.GetMerchByID(p.MerchID)
		if err != nil {
			continue
		}
		inventory = append(inventory, models.InventoryItem{
			Type:     merch.Name,
			Quantity: p.Quantity,
		})
	}

	// Формируем ответ
	response := models.InfoResponse{
		Coins:       user.Coins,
		Inventory:   inventory,
		CoinHistory: coinHistory,
	}

	c.JSON(http.StatusOK, response)
}
