package main

import "fmt"

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type CreditCard struct{}

func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with credit card")
	return true
}

type CheckingAccount struct{}

func (c CheckingAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with checking account")
	return true
}

func main() {
	var option PaymentOption
	option = &CreditCard{}
	option.ProcessPayment(500)

	option = &CheckingAccount{}
	option.ProcessPayment(500)
}
