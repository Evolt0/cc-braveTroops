package status

import (
	"fmt"
)

type HttpError struct {
	// The HTTP status code
	Code int `json:"code"`
	// The error message
	Message string `json:"message"`
}

// HttpError returns the string representation of this error
func (err HttpError) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Message)
}

func New(code int, message string) *HttpError {
	return &HttpError{code, message}
}

func IsError(err error) (*HttpError, bool) {
	if err == nil {
		return nil, false
	}
	if e, ok := err.(HttpError); ok {
		return e.Instantiate(), ok
	}
	e, ok := err.(*HttpError)
	return e, ok
}

func (err HttpError) Instantiate() *HttpError {
	return &HttpError{
		Code:    err.Code,
		Message: err.Message,
	}
}
