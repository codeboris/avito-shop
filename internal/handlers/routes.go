package handlers

import (
	"github.com/codeboris/avito-shop/internal/middleware"
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, userService *service.UserService, merchService *service.MerchService, transactionService *service.TransactionService, authMiddleware *middleware.AuthMiddleware) {
	authHandler := NewAuthHandler(userService)
	infoHandler := NewInfoHandler(userService, merchService, transactionService)
	sendCoinHandler := NewSendCoinHandler(transactionService)
	buyHandler := NewBuyHandler(merchService)

	api := router.Group("/api")
	{
		api.POST("/auth", authHandler.Login)
		api.Use(authMiddleware.Handler())
		api.GET("/info", infoHandler.GetInfo)
		api.POST("/sendCoin", sendCoinHandler.SendCoin)
		api.GET("/buy/:item", buyHandler.BuyItem)
	}
}
