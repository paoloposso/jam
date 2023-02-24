package errors

type NotFoundError struct{}

func CreateNotFoundError() *ValidationError {
	return &ValidationError{}
}

func (e *NotFoundError) Error() string {
	return ""
}
