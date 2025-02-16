package app

import (
	"github.com/codeboris/avito-shop/internal/config"
	"github.com/codeboris/avito-shop/internal/handlers"
	"github.com/codeboris/avito-shop/internal/middleware"
	"github.com/codeboris/avito-shop/internal/repository"
	"github.com/codeboris/avito-shop/internal/service"
	"github.com/codeboris/avito-shop/pkg/jwtutil"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

type App struct {
	router *gin.Engine
	config *config.Config
	db     *sqlx.DB
}

func New() (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Инициализация JWT
	jwtutil.InitJWT(cfg.JWTSecret)

	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		return nil, err
	}

	// Инициализация репозиториев
	userRepo := repository.NewUserRepository(db)
	merchRepo := repository.NewMerchRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	purchaseRepo := repository.NewPurchaseRepository(db)

	// Инициализация сервисов
	userService := service.NewUserService(userRepo)
	merchService := service.NewMerchService(merchRepo, purchaseRepo)
	transactionService := service.NewTransactionService(transactionRepo, userRepo)

	// Инициализация middleware
	authMiddleware := middleware.AuthMiddleware()

	// Инициализация роутера
	router := gin.Default()
	handlers.InitRoutes(router, userService, merchService, transactionService, authMiddleware)

	return &App{
		router: router,
		config: cfg,
		db:     db,
	}, nil
}

func (a *App) Run() error {
	log.Printf("Server is running on port %s", a.config.Server.Port)
	return a.router.Run(":" + a.config.Server.Port)
}
