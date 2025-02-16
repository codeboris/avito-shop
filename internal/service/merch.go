package service

import (
	"errors"
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/codeboris/avito-shop/internal/repository"
)

type MerchService struct {
	merchRepo    *repository.MerchRepository
	purchaseRepo *repository.PurchaseRepository
}

func NewMerchService(merchRepo *repository.MerchRepository, purchaseRepo *repository.PurchaseRepository) *MerchService {
	return &MerchService{
		merchRepo:    merchRepo,
		purchaseRepo: purchaseRepo,
	}
}

// GetMerchItems возвращает список всех товаров
func (s *MerchService) GetMerchItems() ([]models.Merch, error) {
	return s.merchRepo.GetMerchItems()
}

// BuyItem покупает товар
func (s *MerchService) BuyItem(userID, merchID int) error {
	merch, err := s.merchRepo.GetMerchByID(merchID)
	if err != nil {
		return err
	}

	// Проверяем, достаточно ли монет у пользователя
	userRepo := s.purchaseRepo.(*repository.UserRepository) // Приведение типа
	user, err := userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Coins < merch.Price {
		return errors.New("insufficient coins")
	}

	// Обновляем баланс пользователя
	if err := userRepo.UpdateUserCoins(userID, user.Coins-merch.Price); err != nil {
		return err
	}

	// Создаем запись о покупке
	purchase := &models.Purchase{
		UserID:   userID,
		MerchID:  merchID,
		Quantity: 1,
	}
	return s.purchaseRepo.CreatePurchase(purchase)
}

// GetPurchasesByUserID возвращает список покупок пользователя
func (s *MerchService) GetPurchasesByUserID(userID int) ([]models.Purchase, error) {
	return s.purchaseRepo.GetPurchasesByUserID(userID)
}
