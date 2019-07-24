package service

import (
	"context"

	"github.com/sisimogangg/supermarket.discount.api/discount"
	"github.com/sisimogangg/supermarket.discount.api/models"
)

type discountService struct {
	dal discount.DataAccessLayer
}

// NewDicountService creates and returns an instance of discount service
func NewDicountService(repo discount.DataAccessLayer) discount.ServiceLayer {
	return &discountService{dal: repo}
}

func (s *discountService) GetDiscountByProductID(ctx context.Context, productID int32) (*models.Discount, error) {
	disc, err := s.dal.GetDiscountByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return disc, nil
}
