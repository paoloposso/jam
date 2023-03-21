package customerrors

type NotFoundError struct{}

func CreateNotFoundError() *ValidationError {
	return &ValidationError{}
}

func (e *NotFoundError) Error() string {
	return ""
}

type UnauthorizedError struct{}

func CreateUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}

type ValidationError struct {
	s string
}

func CreateValidationError(message string) *ValidationError {
	return &ValidationError{message}
}

func (e *ValidationError) Error() string {
	return e.s
}

type ArgumentError struct {
	s string
}

func CreateArgumentError(message string) *ValidationError {
	return &ValidationError{message}
}

func (e *ArgumentError) Error() string {
	return e.s
}
