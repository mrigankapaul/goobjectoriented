package main

import "fmt"

type PaymentBrokerAccount struct{}

func (p *PaymentBrokerAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment with payment broker")
	return true
}

// client

type PaymentOption interface { // This is being defined in the consuming library.
	ProcessPayment(float32) bool
}

func main() {
	var option PaymentOption

	option = &PaymentBrokerAccount{}

	option.ProcessPayment(500)
}
