package payment

// Option ...
type Option interface {
	ProcessPayment(float32) bool
}
