package dataaccess

import (
	"context"

	"github.com/sisimogangg/supermarket.discount.api/discount"
	"github.com/sisimogangg/supermarket.discount.api/models"
	"github.com/sisimogangg/supermarket.discount.api/utils"
)

type firebaserepo struct{}

var valueD = models.Discount{
	DiscountID:      345,
	Title:           "Buy3Get50%Off",
	Type:            "Value",
	Summary:         "Buy 3 Get 50% Off the third one",
	ProductIDs:      []int32{0, 1},
	AllocationLimit: -1,
}

var productD = models.Discount{
	DiscountID:      346,
	Title:           "Buy2Get1Free",
	Type:            "Product",
	Summary:         "Buy 2 Get 1 Free",
	ProductIDs:      []int32{0, 3},
	AllocationLimit: -1,
}

var appleValueDiscount = models.ValueDiscount{
	Value:            0.50,
	RequiredProducts: 3,
	Discount:         valueD,
}

var coconutProductDiscount = models.ProductDiscount{
	Ratio: struct {
		PreReq   int32 `json:"prereq"`
		Entitled int32 `json:"entitled"`
	}{
		PreReq:   2,
		Entitled: 1,
	},
	Discount: productD,
}

var discounts = [...]discount.Discounter{
	appleValueDiscount,
	coconutProductDiscount,
}

// NewFirebaseRepo creates and returns an instance
func NewFirebaseRepo() discount.DataAccessLayer {
	return &firebaserepo{}
}

func (d *firebaserepo) GetDiscountByProductID(ctx context.Context, productID int32) (*discount.Discounter, error) {
	var discount discount.Discounter
	for _, s := range discounts {
		switch s.(type) {
		case models.ProductDiscount:
			pd := s.(models.ProductDiscount)
			if utils.Contains(productID, pd.ProductIDs) {
				discount = s
			}
		case models.ValueDiscount:
			vd := s.(models.ValueDiscount)

			if utils.Contains(productID, vd.ProductIDs) {
				discount = s
			}

		}
	}

	return &discount, nil
}

func (d *firebaserepo) CheckIfProductIsOnDicount(ctx context.Context, productID int32) (bool, error) {
	isOnDiscount := false
	for _, s := range discounts {
		switch s.(type) {
		case models.ProductDiscount:
			pd := s.(models.ProductDiscount)
			if utils.Contains(productID, pd.ProductIDs) {
				isOnDiscount = true
			}
		case models.ValueDiscount:
			vd := s.(models.ValueDiscount)

			if utils.Contains(productID, vd.ProductIDs) {
				isOnDiscount = true
			}

		}
	}
	return isOnDiscount, nil
}
