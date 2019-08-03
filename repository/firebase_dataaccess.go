package repository

import (
	"context"
	"strings"

	firebase "firebase.google.com/go"
	pb "github.com/sisimogangg/supermarket.discount.api/proto"
)

// Repository represents repo functionality
type Repository interface {
	Get(ctx context.Context, productID string) (*pb.ProductDiscount, error)
	List(ctx context.Context) ([]*pb.ProductDiscount, error)
}

type firebaseRepo struct {
	fb *firebase.App
}

// NewFirebaseRepo defines a constructor for firebaserepo
func NewFirebaseRepo(app *firebase.App) Repository {
	return &firebaseRepo{app}
}

func (f *firebaseRepo) Get(ctx context.Context, productID string) (*pb.ProductDiscount, error) {
	var discount *pb.ProductDiscount

	allDiscounts, err := f.List(ctx) // Temporary Implementation. A better way would be to write a DB Query
	if err != nil {
		return nil, err
	}

	for _, d := range allDiscounts {
		for i := 0; i < len(d.ProductIDs); i++ {
			if strings.Contains(d.ProductIDs[i], productID) {
				discount = d
				break
			}
		}
	}

	return discount, nil
}

func (f *firebaseRepo) List(ctx context.Context) ([]*pb.ProductDiscount, error) {
	client, err := f.fb.Database(ctx)
	if err != nil {
		return nil, err
	}
	discountRef := client.NewRef("discounts")

	pDiscounts := make([]*pb.ProductDiscount, 0, 10)
	var rawResult map[string]pb.ProductDiscount

	err = discountRef.Get(ctx, &rawResult)
	if err != nil {
		return nil, err
	}

	for _, d := range rawResult {
		pDiscounts = append(pDiscounts, &d)
	}

	return pDiscounts, nil
}
