package payment

// CreditCard ...
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// CreateCreditAccount ...
func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard { // constructor
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
		availableCredit: 5000, // this should come from a call to a web service
	}
}

// ProcessPayment ...
func (c CreditCard) ProcessPayment(amount float32) bool {
	return true
}

// AvailableCredit ...
func (c CreditCard) AvailableCredit() float32 {
	return c.availableCredit
}
