package repository

import (
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/jmoiron/sqlx"
)

type MerchRepository struct {
	db *sqlx.DB
}

func NewMerchRepository(db *sqlx.DB) *MerchRepository {
	return &MerchRepository{db: db}
}

// GetMerchItems возвращает список всех товаров
func (r *MerchRepository) GetMerchItems() ([]models.Merch, error) {
	var items []models.Merch
	query := `SELECT id, name, price FROM merch`
	err := r.db.Select(&items, query)
	return items, err
}

// GetMerchByID возвращает товар по ID
func (r *MerchRepository) GetMerchByID(merchID int) (*models.Merch, error) {
	var item models.Merch
	query := `SELECT id, name, price FROM merch WHERE id = $1`
	err := r.db.Get(&item, query, merchID)
	return &item, err
}
