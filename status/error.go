package status

import "fmt"

type ErrorCode string

const (
	Internal        ErrorCode = "internal"
	InvalidArgument ErrorCode = "invalid_argument"
	NotFound        ErrorCode = "not_found"
)

// Error holds an error code, message and error itself.
type Error struct {
	Code     ErrorCode
	Message  string
	Internal error
}

// NewError returns a status error.
func NewError(code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// SetInternal sets an internal error to be logged.
func (e *Error) SetInternal(err error) *Error {
	e.Internal = err
	return e
}

// Error implements error.
func (e *Error) Error() string {
	return fmt.Sprintf("code = %s, message = %s, internal = %s", e.Code, e.Message, e.Internal)
}

// GetCode returns error code.
func (e *Error) GetCode() ErrorCode {
	return e.Code
}

// GetMessage returns error description.
func (e *Error) GetMessage() string {
	return e.Message
}
