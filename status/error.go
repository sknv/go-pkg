package status

import "fmt"

// Error holds an error code, message and error itself.
type Error struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Internal error  `json:"-"`
}

// NewError returns a status error.
func NewError(code int, message string) *Error {
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
