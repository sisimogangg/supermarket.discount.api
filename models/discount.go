package models

// Discount stores discount data
type Discount struct {
	DiscountID      int32   `json:"discountId"`
	Title           string  `json:"title"`
	Type            string  `json:"type"`
	Summary         string  `json:"summary"`
	ProductIDs      []int32 `json:"productIds"`
	AllocationLimit int32   `json:"allocationlimit"`
}
