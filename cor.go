package pract4

import "fmt"

type Receiver struct {
	BankTransfer   bool
	MoneyTransfer  bool
	PayPalTransfer bool
}

type PaymentHandler interface {
	Handle(receiver Receiver)
}

type BankPaymentHandler struct {
	Successor PaymentHandler
}

func (p BankPaymentHandler) Handle(receiver Receiver) {
	if receiver.BankTransfer {
		fmt.Println("Выполняем банковский перевод")
	} else if p.Successor != nil {
		p.Successor.Handle(receiver)
	}
}

type MoneyPaymentHandler struct {
	Successor PaymentHandler
}

func (p MoneyPaymentHandler) Handle(receiver Receiver) {
	if receiver.MoneyTransfer {
		fmt.Println("Выполняем перевод через системы денежных переводов")
	} else if p.Successor != nil {
		p.Successor.Handle(receiver)
	}
}

type PayPalPaymentHandler struct {
	Successor PaymentHandler
}

func (p PayPalPaymentHandler) Handle(receiver Receiver) {
	if receiver.MoneyTransfer {
		fmt.Println("Выполняем перевод через PayPal")
	} else if p.Successor != nil {
		p.Successor.Handle(receiver)
	}
}
