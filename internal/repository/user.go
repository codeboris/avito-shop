package repository

import (
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser создает нового пользователя
func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password, coins) VALUES (:username, :password, :coins)`
	_, err := r.db.NamedExec(query, user)
	return err
}

// GetUserByUsername возвращает пользователя по имени
func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, coins FROM users WHERE username = $1`
	err := r.db.Get(&user, query, username)
	return &user, err
}

// GetUserByID возвращает пользователя по ID
func (r *UserRepository) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, coins FROM users WHERE id = $1`
	err := r.db.Get(&user, query, userID)
	return &user, err
}

// UpdateUserCoins обновляет количество монет пользователя
func (r *UserRepository) UpdateUserCoins(userID, coins int) error {
	query := `UPDATE users SET coins = $1 WHERE id = $2`
	_, err := r.db.Exec(query, coins, userID)
	return err
}
