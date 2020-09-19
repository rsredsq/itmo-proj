package pract4

import "testing"

func TestCor(t *testing.T) {
	receiver := Receiver{false, true, true}

	bankPaymentHandler := BankPaymentHandler{}
	moneyPaymentHandler := MoneyPaymentHandler{}
	paypalPaymentHandler := PayPalPaymentHandler{}

	bankPaymentHandler.Successor = &paypalPaymentHandler
	paypalPaymentHandler.Successor = &moneyPaymentHandler

	bankPaymentHandler.Handle(receiver)
}
