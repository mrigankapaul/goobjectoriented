package main

import "fmt"

// Account ...
type Account struct{}

// AvailableFunds ...
func (a *Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds")
	return 0
}

// ProcessPayment ...
func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

// CreditAccount ...
type CreditAccount struct {
	Account // type embedding
}

func main() {
	ca := &CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)

}
