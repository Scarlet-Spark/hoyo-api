package errors

import "fmt"

// HoYoLab OK status.
const HoyoStatusOK = 0

// HoYoLab custom error handling.
// HoYoLab endpoints return their own special error codes within the response body, despite the response status being 200 OK.
type HoyoError struct {
	RetCode int
	Message string
	Exists  bool
	error
}

// Creates a new HoYoLab error instance.
func NewHoyoError(retcode int, message string) HoyoError {
	hoyoError := HoyoError{
		RetCode: retcode,
		Message: message,
		Exists:  true,
		error:   fmt.Errorf("%d | %s", retcode, message),
	}

	return hoyoError
}
