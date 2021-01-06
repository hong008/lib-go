package errors

import (
	"fmt"
)

type Error interface {
	Error() string
	Code() int
	Message() string
}

type myError struct {
	code    int
	message string
}

func NewError(code int, message string) Error {
	return &myError{
		code:    code,
		message: message,
	}
}

func (m *myError) Error() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("Code: %d, Message: %s", m.code, m.message)
}

func (m *myError) Code() int {
	if m != nil {
		return m.code
	}
	return 0
}

func (m *myError) Message() string {
	if m != nil {
		return m.message
	}
	return ""
}
