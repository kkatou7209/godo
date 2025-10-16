package validation

var (
	ErrEmailAlreadyExists = NewValidationError("email already exists")
	ErrUserNotFound = NewValidationError("user not found")
	ErrInvalidUser = NewValidationError("invalid user")
	ErrTodoNotDound = NewValidationError("todo not found")
	ErrInvalidPassword = NewValidationError("invalid password")
	ErrInvalidTodoInput = NewValidationError("invalid todo input")
)

type ValidationError struct {
	msg string
}

func NewValidationError(msg string) *ValidationError {
	return &ValidationError{msg}
}

func (e *ValidationError) Error() string {
	return e.msg
}