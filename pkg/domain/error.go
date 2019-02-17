package domain

// Error is a custom error that helps with differentiating between internal errors
// and user specific errors
type Error struct {
	code    string
	message string
}

// Code will return the error code
func (e Error) Code() string {
	return e.code
}

// Error is used to conform with the error interface
func (e Error) Error() string {
	return e.message
}

// NewError will create a new custom error
func NewError(code, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}
