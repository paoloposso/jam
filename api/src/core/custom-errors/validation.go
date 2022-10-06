package customerrors

type ValidationError struct {
	s string
}

func NewValidationError(text string) *ValidationError {
	return &ValidationError{text}
}

func (e *ValidationError) Error() string {
	return e.s
}
