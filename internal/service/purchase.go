package service

import (
	"context"
	"errors"
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/codeboris/avito-shop/internal/repository"
)

type PurchaseService struct {
	purchaseRepo *repository.PurchaseRepository
	userRepo     *repository.UserRepository
	merchRepo    *repository.MerchRepository
}

func NewPurchaseService(purchaseRepo *repository.PurchaseRepository, userRepo *repository.UserRepository, merchRepo *repository.MerchRepository) *PurchaseService {
	return &PurchaseService{purchaseRepo: purchaseRepo, userRepo: userRepo, merchRepo: merchRepo}
}

func (s *PurchaseService) BuyMerch(ctx context.Context, userID int, merchName string) error {
	merch, err := s.merchRepo.GetMerchByName(ctx, merchName)
	if err != nil {
		return err
	}
	if merch == nil {
		return errors.New("merch not found")
	}

	user, err := s.userRepo.GetUserByUsername(ctx, "")
	if err != nil {
		return err
	}

	if user.Coins < merch.Price {
		return errors.New("insufficient coins")
	}

	if err := s.userRepo.UpdateUserCoins(ctx, userID, user.Coins-merch.Price); err != nil {
		return err
	}

	purchase := &models.Purchase{
		UserID:   userID,
		MerchID:  merch.ID,
		Quantity: 1,
	}

	return s.purchaseRepo.CreatePurchase(ctx, purchase)
}
