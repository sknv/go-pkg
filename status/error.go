package status

import (
	"fmt"
)

type ErrorCode int

// Error holds an error code, message and error itself.
type Error struct {
	Code     ErrorCode `json:"code,omitempty"`
	Message  string    `json:"message,omitempty"`
	Internal error     `json:"-"`
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

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d, msg = %s, err = %s", e.Code, e.Message, e.Internal)
}
