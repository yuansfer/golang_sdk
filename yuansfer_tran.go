package yuansfer

// Yuansfer payment method
type Yuansfer interface {
	PostToYuansfer() (string, error)
}
