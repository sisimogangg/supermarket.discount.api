package service

import (
	"context"
	"time"

	"github.com/sisimogangg/supermarket.discount.api/discount"
	"github.com/sisimogangg/supermarket.discount.api/models"
)

type discountService struct {
	dal     discount.DataAccessLayer
	timeOut time.Duration
}

// NewDicountService creates and returns an instance of discount service
func NewDicountService(repo discount.DataAccessLayer, timeout time.Duration) discount.ServiceLayer {
	return &discountService{dal: repo, timeOut: timeout}
}

func (s *discountService) GetDiscountByProductID(ctx context.Context, productID int32) (*models.Discount, error) {
	disc, err := s.dal.GetDiscountByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return disc, nil
}
