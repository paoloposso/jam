package errors

type ValidationError struct {
	s string
}

func CreateValidationError(message string) *ValidationError {
	return &ValidationError{message}
}

func (e *ValidationError) Error() string {
	return e.s
}
