package main

import "fmt"

// CreditAccount ...
type CreditAccount struct{}

// AvailableFunds ...
func (c *CreditAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit funds")
	return 250
}

// CheckingAccount ...
type CheckingAccount struct{}

// AvailableFunds ...
func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds")
	return 125
}

// HybridAccount ...
type HybridAccount struct {
	creditAccount   CreditAccount
	checkingAccount CheckingAccount
}

// AvailableFunds ...
func (h *HybridAccount) AvailableFunds() float32 {
	return h.creditAccount.AvailableFunds() + h.checkingAccount.AvailableFunds()
}

func main() {
	ha := &HybridAccount{}
	fmt.Println(ha.AvailableFunds())
}
