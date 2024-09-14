package factory

import (
	"fmt"
)

// The `PaymentMethod` interface defines a contract for payment methods. Here's what each method does:
// - `Pay(amount float32) string`: This method takes an `amount` of type `float32` as input and returns a `string` representing the payment method used.
// In other words, any type that implements the `PaymentMethod` interface must provide an implementation for the `Pay` method.
type PaymentMethod interface {
	Pay(amount float32) string
}

// PaymentMethods is an enumeration of possible payment methods.
type PaymentMethods int

// The enumeration values represent the different payment methods.
const (
	Cash PaymentMethods = iota + 1
	DebitCard
)

// GetPaymentMethod returns a PaymentMethod instance for the given PaymentMethods
// value, or an error if the value is not recognized.
func GetPaymentMethod(m PaymentMethods) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		// return new(DebitCardPM), nil
		return new(CreditCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}

// CashPM represents the cash payment method.
type CashPM struct{}

// Pay takes an amount of type float32 as input and returns a string representing the
// payment method used. In this case, it returns a string indicating that the payment
// was made using cash, with the amount formatted to two decimal places.
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// DebitCardPM represents the debit card payment method.
type DebitCardPM struct{}

// Pay takes an amount of type float32 as input and returns a string representing the
// payment method used. In this case, it returns a string indicating that the payment
// was made using debit card, with the amount formatted to two decimal places.
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

// CreditCardPM represents the credit card payment method.
type CreditCardPM struct{}

// Pay takes an amount of type float32 as input and returns a string representing the
// payment method used. In this case, it returns a string indicating that the payment
// was made using credit card, with the amount formatted to two decimal places.
func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using new credit card implementation\n", amount)
	// return fmt.Sprintf("%#0.2f paid using debit card (new)\n", amount)
}
