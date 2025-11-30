package srp

// SRP: Single Responsibility Principle

// SRP - Violated
type OrderProcessorViolation struct{}

func (os *OrderProcessorViolation) ProcessOrder() {
	// calculate total
	os.CalculateTotal()

	// process payment
	os.ProcessPayment()

	// send confirmation email
	os.SendConfirmationEmail()
}

func (os *OrderProcessorViolation) CalculateTotal()        {}
func (os *OrderProcessorViolation) ProcessPayment()        {}
func (os *OrderProcessorViolation) SendConfirmationEmail() {}

// SRP - Followed
// price calculation
type PriceCalculator struct{}

func (pc *PriceCalculator) CalculateTotal() {}

// payment processing
type PaymentProcessor struct{}

func (pp *PaymentProcessor) ProcessPayment() {}

// email notification
type EmailNotifier struct{}

func (en *EmailNotifier) SendConfirmationEmail() {}

type OrderProcessor struct {
	priceCalculator  *PriceCalculator
	paymentProcessor *PaymentProcessor
	emailNotifier    *EmailNotifier
}

func (op *OrderProcessor) ProcessOrder() {
	// calculate total
	op.priceCalculator.CalculateTotal()

	// process payment
	op.paymentProcessor.ProcessPayment()

	// send confirmation email
	op.emailNotifier.SendConfirmationEmail()
}
