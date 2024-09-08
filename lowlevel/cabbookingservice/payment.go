package main

type paymentType string

const (
	upi  paymentType = "UPI"
	cash paymentType = "CASH"
)

type paymentMethods interface {
	pay(amount float64) string
}

func newPaymentFactory(paymentType paymentType) paymentMethods {
	switch paymentType {
	case upi:
		return &upiPayment{}
	case cash:
		return &cashPayment{}
	}

	// Default
	return &cashPayment{}
}

type upiPayment struct{}

func (u *upiPayment) pay(amount float64) string {
	return "Paid using UPI"
}

type cashPayment struct{}

func (c *cashPayment) pay(amount float64) string {
	return "Paid using CASH"
}
