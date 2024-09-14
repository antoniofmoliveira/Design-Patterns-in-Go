package factory

import (
	"strings"
	"testing"
)

// This is a Go test function that verifies the behavior of the `GetPaymentMethod` function when retrieving the "Cash" payment method. It checks that:

// 1. The "Cash" payment method exists (i.e., `GetPaymentMethod` returns no error).
// 2. When making a payment of $10.30 using the "Cash" method, the resulting message contains the phrase "paid using cash".

// If either of these conditions is not met, the test fails and logs an error message.
func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

// This is a Go test function that verifies the behavior of the `GetPaymentMethod` function when retrieving the "DebitCard" payment method. It checks that:
// 1. The "DebitCard" payment method exists (i.e., `GetPaymentMethod` returns no error).
// 2. When making a payment of $22.30 using the "DebitCard" method, the resulting message contains the phrase "paid using debit card".
// If either of these conditions is not met, the test fails and logs an error message.
func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Error("A payment method of type 'DebitCard' must exist")
	}
	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

// This is a Go test function that verifies the behavior of the `GetPaymentMethod` function when retrieving a non-existent payment method. It checks that:
// 1. When retrieving a payment method with ID 20, `GetPaymentMethod` returns an error.
// If this condition is not met, the test fails and logs an error message.
func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
