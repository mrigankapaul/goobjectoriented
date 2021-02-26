package main

import (
	"fmt"
	"strings"

	"mrigankapaul/payment"
)

func main() {
	const amount = 500

	fmt.Println("Paying with cash")
	cash := &payment.Cash{}
	cash.ProcessPayment(amount)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	credit := payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	fmt.Println("Paying with credit card")
	fmt.Printf("Initial balance: $%.2f\n", credit.AvailableCredit())
	credit.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f\n", credit.AvailableCredit())
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	checking := payment.CreateCheckingAccount(
		"Debra Ingram",
		"01010011",
		"0551005115100")

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance())
	checking.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f\n", checking.Balance())

	fmt.Println("Hmm, not enough in the account. Nothing we can do!")
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")
}
