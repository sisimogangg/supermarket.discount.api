package service

import (
	"context"
	"time"

	pb "github.com/sisimogangg/supermarket.discount.api/proto"
	"github.com/sisimogangg/supermarket.discount.api/repository"
)

type discountService struct {
	repo    repository.Repository
	timeOut time.Duration
}

// NewDicountService creates and returns an instance of discount service
func NewDicountService(repo repository.Repository, timeout time.Duration) pb.DiscountServiceHandler {
	return &discountService{repo, timeout}
}

func (s *discountService) Get(ctx context.Context, req *pb.GetRequest, res *pb.ProductDiscount) error {
	discount, err := s.repo.Get(ctx, req.ProductID)
	if err != nil {
		return err
	}

	if discount != nil {
		res.DiscountID = discount.DiscountID
		res.Title = discount.Title
		res.Summary = discount.Summary
		res.Type = discount.Type
		res.Allocationlimit = discount.Allocationlimit
		res.DiscountRatio = discount.DiscountRatio
		res.DiscountValue = discount.DiscountValue
		res.ProductIDs = discount.ProductIDs
	}

	return nil
}

func (s *discountService) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	discounts, err := s.repo.List(ctx)
	if err != nil {
		return err
	}

	res.Discounts = discounts
	return nil
}
