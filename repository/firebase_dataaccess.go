package dataaccess

import (
	"context"

	pb "github.com/sisimogangg/supermarket.discount.api/proto"
)

// Repository represents repo functionality
type Repository interface {
	Get(ctx context.Context, productID string) (*pb.ProductDiscount, error)
	List(ctx context.Context) ([]*pb.ProductDiscount, error)
}

type firebaserepo struct{}

var discountValue = pb.DiscountValue{
	Value:            0.50,
	RequiredProducts: 3,
}

var discountRatio = pb.DiscountRatio{
	PreRequisite: 2,
	Entitled:     1,
}

var valueD = pb.ProductDiscount{
	DiscountID:      "345",
	Title:           "Buy3Get50%Off",
	Summary:         "Buy 3 Get 50% Off the third one",
	Type:            "Value",
	DiscountValue:   &discountValue,
	ProductIDs:      []string{"100", "200"},
	Allocationlimit: -1,
}

var productD = pb.ProductDiscount{
	DiscountID:      "346",
	Title:           "Buy2Get1Free",
	Summary:         "Buy 2 Get 1 Free",
	Type:            "Product",
	DiscountRatio:   &discountRatio,
	ProductIDs:      []string{"100", "200"},
	Allocationlimit: -1,
}

// NewFirebaseRepo creates and returns an instance
func NewFirebaseRepo() Repository {
	return &firebaserepo{}
}

func (d *firebaserepo) Get(ctx context.Context, productID int32) (*pb.ProductDiscount, error) {
	var discount pb.ProductDiscount

	return &discount, nil
}

func (d *firebaserepo) List(ctx context.Context) ([]*pb.ProductDiscount, error) {
	pDiscounts := make([]*pb.ProductDiscount, 0)
	return pDiscounts, nil
}
