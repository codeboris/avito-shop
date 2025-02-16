package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, coins, created_at, updated_at 
	          FROM users 
	          WHERE username = $1`

	err := r.db.GetContext(ctx, &user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, username string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("default_password"), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Username:  username,
		Password:  string(hashedPassword),
		Coins:     1000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `INSERT INTO users (username, password, coins, created_at, updated_at) 
	          VALUES (:username, :password, :coins, :created_at, :updated_at) 
	          RETURNING id`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &newUser.ID, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *UserRepository) GetOrCreateUser(ctx context.Context, username string) (*models.User, error) {
	user, err := r.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = r.CreateUser(ctx, username)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
