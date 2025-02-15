package db

import (
	"context"
	"fmt"
	"github.com/codeboris/avito-shop/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func buildDSN(cfg config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name)
}

func NewPostgresDB(cfg config.DatabaseConfig) (*sqlx.DB, error) {
	dsn := buildDSN(cfg)
	dbApp, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := dbApp.PingContext(ctx); err != nil {
		dbApp.Close()
		return nil, err
	}

	return dbApp, nil
}
