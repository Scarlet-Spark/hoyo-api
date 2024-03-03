package handler

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Scarlet-Spark/hoyo-api/internal/errors"
	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

// Default HTTP client timeout duration.
const CLIENT_DEFAULT_TIMEOUT = 10 * time.Second

// Base HTTP request handler for sending requests to HoYoLab endpoints.
type Handler struct {
	client http.Client
	cache  middleware.Cache
}

// Creates a new HTTP request handler instance.
func NewHandler() *Handler {
	handler := &Handler{http.Client{Timeout: CLIENT_DEFAULT_TIMEOUT}, *middleware.NewCache()}
	return handler
}

// Sends a HTTP request.
// Returns a generic map from the unmarshalled response.
// Specific retcode errors are handled by their respective clients.
func (handler *Handler) Send(request *Request) (data map[string]interface{}, clientError errors.HandlerError) {
	// Build HTTP request.
	httpRequest, clientError := handler.createHttpRequest(request)

	if clientError.Exists {
		return
	}

	// Send HTTP request.
	response, err := handler.client.Do(httpRequest)
	if err != nil {
		return nil,
			errors.NewHandlerError(
				errors.HTTP_REQUEST_SEND_ERROR.String(),
				fmt.Sprintf("URL: %s, Error: %s", request.endpoint, err.Error()),
			)
	}

	defer response.Body.Close()

	// Parse response body into readable format.
	body, clientError := handler.parseResponse(response)
	if clientError.Exists {
		return nil, clientError
	}

	if response.StatusCode != http.StatusOK {
		return nil,
			errors.NewHandlerError(
				response.Status,
				fmt.Sprintf("URL: %s, Status Code: %d, Error: %+v", request.endpoint, response.StatusCode, string(body)),
			)
	}

	// Return JSON marshalled response.
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil,
			errors.NewHandlerError(
				errors.JSON_DESERIALIZATION_ERROR.String(),
				fmt.Sprintf("Body: %s, Error: %s", body, err.Error()),
			)
	}

	return data, clientError
}

// Create HTTP request structure.
func (handler *Handler) createHttpRequest(request *Request) (httpRequest *http.Request, clientError errors.HandlerError) {
	var body io.Reader

	// JSON marshal request body.
	if request.body != nil {
		jsonData, err := json.Marshal(request.body)

		if err != nil {
			return nil,
				errors.NewHandlerError(
					errors.JSON_SERIALIZATION_ERROR.String(),
					fmt.Sprintf("Request Body: %+v, Error: %s", request.body, err.Error()),
				)
		}

		body = bytes.NewBuffer(jsonData)
	}

	// Create HTTP request.
	httpRequest, err := http.NewRequest(string(request.method), string(request.endpoint), body)

	if err != nil {
		return nil,
			errors.NewHandlerError(
				errors.HTTP_REQUEST_CREATE_ERROR.String(),
				fmt.Sprintf("URL: %s, Error: %s", request.endpoint, err.Error()),
			)
	}

	// Add query parameters to HTTP request.
	query := httpRequest.URL.Query()
	for k, v := range request.params {
		query.Add(k, v)
	}

	httpRequest.URL.RawQuery = query.Encode()

	// Add headers to HTTP request.
	for k, v := range request.headers {
		httpRequest.Header.Set(k, v)
	}

	return httpRequest, clientError
}

// Parse response body by decompressing content according to its encoding.
// HoYoLab endpoints currently support gzip, deflate and brotli.
func (handler *Handler) parseResponse(response *http.Response) (body []byte, clientError errors.HandlerError) {
	var err error
	var reader io.ReadCloser

	switch encoding := response.Header.Get("Content-Encoding"); encoding {
	case string(GZIP):
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil,
				errors.NewHandlerError(
					errors.HTTP_RESPONSE_READ_ERROR.String(),
					fmt.Sprintf("Body: %+v, Error: %s", response.Body, err.Error()),
				)
		}

	case string(DEFLATE):
		reader = flate.NewReader(response.Body)

	default:
		reader = response.Body
	}

	body, err = io.ReadAll(reader)
	reader.Close()

	if err != nil {
		return nil,
			errors.NewHandlerError(
				errors.HTTP_RESPONSE_READ_ERROR.String(),
				fmt.Sprintf("Body: %+v, Error: %s", response.Body, err.Error()),
			)
	}

	return body, clientError
}
