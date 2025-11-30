package ocp

// OCP: Open/Closed Principle

// OCP - Violated
type OrderType string
const (
	RegularOrder OrderType = "regular"
	PremiumOrder OrderType = "premium"
	BlackFriday  OrderType = "black_friday"
)

type PriceCalculator struct{
	OrderType OrderType
}
func (pc *PriceCalculator) CalculateTotal() {
	switch pc.OrderType {
	case RegularOrder:
		// apply regular pricing or no discount
		case PremiumOrder:
		// apply premium member discount
	case BlackFriday:
		// apply black friday discount
	}

	// calculate total price
}


// OCP - Followed
type IDiscountStrategy interface {
	ApplyDiscount(price float64) float64
}

type RegularDiscount struct{}
func (rd *RegularDiscount) ApplyDiscount(price float64) float64 {
	// no discount for regular orders
	return price
}

type PremiumDiscount struct{}
func (pd *PremiumDiscount) ApplyDiscount(price float64) float64 {
	// apply premium member discount
	return price * 0.9 // example: 10% discount
}

type BlackFridayDiscount struct{}
func (bfd *BlackFridayDiscount) ApplyDiscount(price float64) float64 {
	// apply black friday discount
	return price * 0.7 // example: 30% discount
}

type PriceCalculatorOCP struct{
	discountStrategy IDiscountStrategy
}
func (pc *PriceCalculatorOCP) CalculateTotal(price float64) float64 {
	_ = pc.discountStrategy.ApplyDiscount(price)
	// cacculate total price after discount
	return price
}