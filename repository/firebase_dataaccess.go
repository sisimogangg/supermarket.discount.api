package repository

import (
	"context"
	"fmt"

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
	client, err := f.fb.Database(ctx)
	if err != nil {
		return nil, err
	}
	discountRef := client.NewRef(fmt.Sprintf("discounts/%s", productID))

	discount := pb.ProductDiscount{}
	if err := discountRef.Get(ctx, &discount); err != nil {
		return nil, err
	}

	return &discount, nil
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
