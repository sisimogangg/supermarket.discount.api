package models

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
