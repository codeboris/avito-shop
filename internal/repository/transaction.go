package repository

import (
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// CreateTransaction создает новую транзакцию
func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	query := `INSERT INTO transactions (from_user, to_user, amount) VALUES (:from_user, :to_user, :amount)`
	_, err := r.db.NamedExec(query, transaction)
	return err
}

// GetTransactionsByUserID возвращает историю транзакций пользователя
func (r *TransactionRepository) GetTransactionsByUserID(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := `SELECT id, from_user, to_user, amount, created_at FROM transactions WHERE from_user = $1 OR to_user = $1`
	err := r.db.Select(&transactions, query, userID)
	return transactions, err
}
