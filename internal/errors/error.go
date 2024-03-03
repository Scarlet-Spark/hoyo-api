package errors

import (
	"fmt"
)

// Custom error handling.
type Error struct {
	ErrorCode    ErrorCode
	ErrorMessage string
	error
}

// Creates a new error instance.
func NewError(errorCode ErrorCode, errorMessage string) Error {
	hoyoApiErr := Error{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		error:        fmt.Errorf("%s | %s", errorCode, errorMessage),
	}

	return hoyoApiErr
}
