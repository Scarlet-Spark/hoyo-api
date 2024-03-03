package errors

// Custom error codes for Handler.
type ErrorCode string

const (
	// Internal error codes.
	JSON_SERIALIZATION_ERROR   ErrorCode = "JSON_SERIALIZATION_ERROR"
	JSON_DESERIALIZATION_ERROR ErrorCode = "JSON_DESERIALIZATION_ERROR"
	HTTP_REQUEST_CREATE_ERROR  ErrorCode = "HTTP_REQUEST_CREATE_ERROR"
	HTTP_REQUEST_SEND_ERROR    ErrorCode = "HTTP_REQUEST_SEND_ERROR"
	HTTP_RESPONSE_STATUS_ERROR ErrorCode = "HTTP_RESPONSE_STATUS_ERROR"
	HTTP_RESPONSE_READ_ERROR   ErrorCode = "HTTP_RESPONSE_READ_ERROR"
)

// Error code stringer.
func (errorCode ErrorCode) String() string {
	return string(errorCode)
}
