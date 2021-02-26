package payment

// Cash ...
type Cash struct{}

// CreateCashAccount ...
func CreateCashAccount() *Cash { // constructor
	return &Cash{}
}

// ProcessPayment ...
func (c Cash) ProcessPayment(amount float32) bool {
	return true
}
