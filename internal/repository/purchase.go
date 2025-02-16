package repository

import (
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/jmoiron/sqlx"
)

type PurchaseRepository struct {
	db *sqlx.DB
}

func NewPurchaseRepository(db *sqlx.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

// CreatePurchase создает запись о покупке
func (r *PurchaseRepository) CreatePurchase(purchase *models.Purchase) error {
	query := `INSERT INTO purchases (user_id, merch_id, quantity) VALUES (:user_id, :merch_id, :quantity)`
	_, err := r.db.NamedExec(query, purchase)
	return err
}

// GetPurchasesByUserID возвращает список покупок пользователя
func (r *PurchaseRepository) GetPurchasesByUserID(userID int) ([]models.Purchase, error) {
	var purchases []models.Purchase
	query := `SELECT id, user_id, merch_id, quantity, created_at FROM purchases WHERE user_id = $1`
	err := r.db.Select(&purchases, query, userID)
	return purchases, err
}
