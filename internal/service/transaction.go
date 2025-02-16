package service

import (
	"errors"
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/codeboris/avito-shop/internal/repository"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
	userRepo        *repository.UserRepository
}

func NewTransactionService(transactionRepo *repository.TransactionRepository, userRepo *repository.UserRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
	}
}

// SendCoins отправляет монеты от одного пользователя другому
func (s *TransactionService) SendCoins(fromUserID, toUserID, amount int) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	fromUser, err := s.userRepo.GetUserByID(fromUserID)
	if err != nil {
		return err
	}

	if fromUser.Coins < amount {
		return errors.New("insufficient coins")
	}

	toUser, err := s.userRepo.GetUserByID(toUserID)
	if err != nil {
		return err
	}

	// Обновляем баланс отправителя и получателя
	if err := s.userRepo.UpdateUserCoins(fromUserID, fromUser.Coins-amount); err != nil {
		return err
	}
	if err := s.userRepo.UpdateUserCoins(toUserID, toUser.Coins+amount); err != nil {
		return err
	}

	// Создаем запись о транзакции
	transaction := &models.Transaction{
		FromUser: fromUserID,
		ToUser:   toUserID,
		Amount:   amount,
	}
	return s.transactionRepo.CreateTransaction(transaction)
}

// GetTransactionHistory возвращает историю транзакций пользователя
func (s *TransactionService) GetTransactionHistory(userID int) ([]models.Transaction, error) {
	return s.transactionRepo.GetTransactionsByUserID(userID)
}
