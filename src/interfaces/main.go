package main

import "interface/payment"

func main() {

	var option payment.PaymentOption

	option = payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	option.ProcessPayment(500) // message passing

	option = payment.CreateCashAccount()

	option.ProcessPayment(500) // message passing
}
