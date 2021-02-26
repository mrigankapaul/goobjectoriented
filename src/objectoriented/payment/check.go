package payment

// CheckingAccount ...
type CheckingAccount struct {
	accountOwner  string
	routingNumber string
	accountNumber string
	balance       float32
}

// CreateCheckingAccount ...
func CreateCheckingAccount(accountOwner, routingNumber, accountNumber string) *CheckingAccount {
	return &CheckingAccount{
		accountOwner:  accountOwner,
		routingNumber: routingNumber,
		accountNumber: accountNumber,
		balance:       250, // this should come from a call to a web service
	}
}

// ProcessPayment ...
func (c CheckingAccount) ProcessPayment(amount float32) bool {
	return true
}

// Balance ...
func (c CheckingAccount) Balance() float32 {
	return c.balance
}
