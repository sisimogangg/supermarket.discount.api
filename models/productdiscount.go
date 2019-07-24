package models

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
