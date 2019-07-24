package dataaccess

import (
	"context"

	"github.com/sisimogangg/supermarket.discount.api/discount"
	"github.com/sisimogangg/supermarket.discount.api/models"
)

type firebaserepo struct{}

var discounts = [...]models.Discount{
	models.Discount{
		DiscountID:      345,
		Title:           "Buy3Get50%Off",
		Type:            "Value",
		Summary:         "Buy 3 Get 50% Off the third one",
		ProductIDs:      []int32{0, 3},
		Value:           "-1.00",
		AllocationLimit: -1,
	},
	models.Discount{
		DiscountID: 346,
		Title:      "Buy2Get1Free",
		Type:       "Product",
		Summary:    "Buy 2 Get 1 Free",
		ProductIDs: []int32{0, 3},
		Ratio: struct {
			PreReq   int32
			Entitled int32
		}{
			PreReq:   2,
			Entitled: 1,
		},
		AllocationLimit: -1,
	},
}

// NewFirebaseRepo creates and returns an instance
func NewFirebaseRepo() discount.DataAccessLayer {
	return &firebaserepo{}
}

func (d *firebaserepo) GetDiscountByProductID(ctx context.Context, productID int32) (*models.Discount, error) {
	discount := models.Discount{}
	for _, s := range discounts {
		for _, p := range s.ProductIDs {
			if p == productID {
				discount = s
				break
			}
		}
	}

	if discount.DiscountID == 0 {
		return nil, nil
	}

	return &discount, nil
}
