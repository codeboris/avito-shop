package app

import (
	"github.com/codeboris/avito-shop/internal/config"
	"github.com/codeboris/avito-shop/internal/db"
	"github.com/codeboris/avito-shop/internal/handler"
	"github.com/codeboris/avito-shop/internal/server"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
)

type App struct {
	Router *mux.Router
	Config *config.Config
	Server *server.Server
	DB     *sqlx.DB
}

func New() (*App, error) {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config file: %s", err)
		return nil, err
	}

	dbConn, err := db.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
		return nil, err
	}

	router := mux.NewRouter()
	handler.InitHandlers(router)

	srv := server.New(cfg.Server.Port, router)

	return &App{
		Router: router,
		Config: cfg,
		Server: srv,
		DB:     dbConn,
	}, nil
}

func (a *App) Run() error {
	return a.Server.Run()
}
