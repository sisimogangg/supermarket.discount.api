package discount

// Discounter defines behaviour of product discounter
type Discounter interface {
	CalculateDiscount() float32
}
