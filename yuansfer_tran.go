package yuansfer

type Yuansfer interface {
	PostToYuansfer(token string) (string, error)
}
