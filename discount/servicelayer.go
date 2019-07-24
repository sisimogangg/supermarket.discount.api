package discount

import (
	"context"

	"github.com/sisimogangg/supermarket.discount.api/models"
)

// ServiceLayer defines expected service layer behavour
type ServiceLayer interface {
	GetDiscountByProductID(ctx context.Context, productID int32) (*models.Discounter, error)
}
