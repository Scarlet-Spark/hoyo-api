package errors

import "net/http"

const (
	// Internal error codes.
	JSON_SERIALIZATION_ERROR = iota
	JSON_DESERIALIZATION_ERROR
	HTTP_REQUEST_CREATE_ERROR
	HTTP_REQUEST_SEND_ERROR
	HTTP_RESPONSE_STATUS_ERROR
	HTTP_RESPONSE_READ_ERROR
	HTTP_RESPONSE_CLOSE_ERROR
	REGION_SERVER_CODE_ERROR
)

// Internal error code text.
var errorCodeText = map[int]string{
	JSON_SERIALIZATION_ERROR:   "JSON Serialization Error",
	JSON_DESERIALIZATION_ERROR: "JSON Deserialization Error",
	HTTP_REQUEST_CREATE_ERROR:  "HTTP Request Create Error",
	HTTP_REQUEST_SEND_ERROR:    "HTTP Request Send Error",
	HTTP_RESPONSE_STATUS_ERROR: "HTTP Response Status Error",
	HTTP_RESPONSE_READ_ERROR:   "HTTP Response Read Error",
	HTTP_RESPONSE_CLOSE_ERROR:  "HTTP Response Close Error",
	REGION_SERVER_CODE_ERROR:   "Region Server Code Error",
}

// Returns a text for the internal error code.
// Returns an empty string if the code is unknown.
func ErrorCodeText(code int) string {
	text, exists := errorCodeText[code]
	if !exists {
		text = http.StatusText(code)
	}

	return text
}
