package utils

import pb "github.com/sisimogangg/supermarket.discount.api/proto"

// ProductDiscounts store seeding data
var ProductDiscounts = []*pb.ProductDiscount{
	&pb.ProductDiscount{
		DiscountID: "345",
		Title:      "Buy3Get50%Off",
		Summary:    "Buy 3 Get 50% Off the third one",
		Type:       "Value",
		DiscountValue: &pb.DiscountValue{
			Value:            0.50,
			RequiredProducts: 3,
		},
		ProductIDs:      []string{"100", "200"},
		Allocationlimit: -1,
	},
	&pb.ProductDiscount{
		DiscountID: "346",
		Title:      "Buy2Get1Free",
		Summary:    "Buy 2 Get 1 Free",
		Type:       "Product",
		DiscountRatio: &pb.DiscountRatio{
			PreRequisite: 2,
			Entitled:     1,
		},
		ProductIDs:      []string{"100", "200"},
		Allocationlimit: -1,
	},
}
