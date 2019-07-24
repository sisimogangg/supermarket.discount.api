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

//ValueDiscount stores the value discount
type ValueDiscount struct {
	Value            float32 `json:"value"`
	RequiredProducts int32   `json:"requiredproducts"`
	Discount
}

//CalculateDiscount returns a discount
func (d ValueDiscount) CalculateDiscount() float32 {
	return 0.0
}

//ProductDiscount stores product discount
type ProductDiscount struct {
	Ratio struct {
		PreReq   int32 `json:"prereq"`
		Entitled int32 `json:"entitled"`
	}
	Discount
}

//CalculateDiscount returns a discount
func (d ProductDiscount) CalculateDiscount() float32 {
	return 0.0
}

// Discounter defines behaviour of product discounter
type Discounter interface {
	CalculateDiscount() float32
}
