package discount

import (
	"context"
)

// ServiceLayer defines expected service layer behavour
type ServiceLayer interface {
	GetDiscountByProductID(ctx context.Context, productID int32) (*Discounter, error)
	CheckIfProductIsOnDicount(ctx context.Context, productID int32) (bool, error)
}
