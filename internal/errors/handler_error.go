package errors

import (
	"fmt"
)

// Handler custom error handling.
type HandlerError struct {
	ErrorCode    string
	ErrorMessage string
	Exists       bool
	error
}

// Creates a new handler error instance.
func NewHandlerError(errorCode string, errorMessage string) HandlerError {
	clientError := HandlerError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		Exists:       true,
		error:        fmt.Errorf("%s | %s", errorCode, errorMessage),
	}

	return clientError
}
