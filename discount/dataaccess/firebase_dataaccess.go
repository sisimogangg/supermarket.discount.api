package dataaccess

import (
	"context"
	"github.com/sisimogangg/supermarket.discount.api/discount"
	"github.com/sisimogangg/supermarket.discount.api/models"
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
		PreReq: 2,
		Entitled: 1,
	},
	Discount: productD,
}

var discounts = [...]models.Discounter{
	appleValueDiscount,
	coconutProductDiscount,
}

// NewFirebaseRepo creates and returns an instance
func NewFirebaseRepo() discount.DataAccessLayer {
	return &firebaserepo{}
}

func (d *firebaserepo) GetDiscountByProductID(ctx context.Context, productID int32) (*models.Discounter, error) {
	var discount models.Discounter
	for _, s := range discounts {
		switch s.(type) {
		case models.ProductDiscount:
			pd := s.(models.ProductDiscount)
			for _, r := range pd.ProductIDs {
				if r == productID {
					discount = s
				}
			}
		case models.ValueDiscount:
			vd := s.(models.ValueDiscount)
			for _, id := range vd.ProductIDs {
				if id == productID {
					discount = s
				}
			}

		}
	}

	return &discount, nil
}
