package errors

import (
	"fmt"
)

// Custom error for handling internal errors and HoYoLab endpoint errors.
// Errors from HoYoLab endpoints have special error codes within the response body after returning a 200 OK response.
type Error struct {
	ErrorCode    int
	ErrorMessage string
}

// Constructor.
func NewError(errorCode int, errorMessage string) error {
	err := Error{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}

	return err
}

// Returns custom error message.
// Implements from built-in error interface.
func (handler Error) Error() string {
	return fmt.Sprintf("%d %s\n%s", handler.ErrorCode, ErrorCodeText(handler.ErrorCode), handler.ErrorMessage)
}