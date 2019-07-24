package discount

import (
	"context"

	"github.com/sisimogangg/supermarket.discount.api/models"
)

// DataAccessLayer defines behaviour of repository
type DataAccessLayer interface {
	GetDiscountByProductID(ctx context.Context, productID int32) (*models.Discount, error)
}
