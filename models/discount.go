package models

// Discount stores discount data
type Discount struct {
	DiscountID int32
	Title      string
	Type       string
	Summary    string
	ProductIDs []int32
	Value      string
	Ratio      struct {
		PreReq   int32
		Entitled int32
	}
	AllocationLimit int32
}
