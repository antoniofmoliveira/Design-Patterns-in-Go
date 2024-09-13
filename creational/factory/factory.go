package factory

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

type PaymentMethods int

const (
	Cash PaymentMethods = iota + 1
	DebitCard
)

// const (
// 	Cash      = 1
// 	DebitCard = 2
// )

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

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// type DebitCardPM struct{}

// func (c *DebitCardPM) Pay(amount float32) string {
// 	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
// }

type CreditCardPM struct{}

func (d *CreditCardPM) Pay(amount float32) string {
	// return fmt.Sprintf("%#0.2f paid using new credit card implementation\n", amount)
	return fmt.Sprintf("%#0.2f paid using debit card (new)\n", amount)
}
