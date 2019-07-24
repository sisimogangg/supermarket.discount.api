package discount

import (
	"context"
)

// DataAccessLayer defines behaviour of repository
type DataAccessLayer interface {
	GetDiscountByProductID(ctx context.Context, productID int32) (*Discounter, error)
	CheckIfProductIsOnDicount(ctx context.Context, productID int32) (bool, error)
}
