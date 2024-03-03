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

const (
	// Default HTTP client timeout duration.
	CLIENT_DEFAULT_TIMEOUT = 10 * time.Second
)

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
func (handler *Handler) Send(request *Request) (map[string]interface{}, error) {
	// Build HTTP request.
	httpRequest, err := handler.createHttpRequest(request)

	if err != nil {
		return nil, err
	}

	// Send HTTP request.
	response, err := handler.client.Do(httpRequest)
	if err != nil {
		return nil, errors.NewError(errors.API_REQUEST_SEND_ERROR, fmt.Sprintf("URL: %s, Error: %s", request.endpoint, err.Error()))
	}

	defer response.Body.Close()

	// Parse response body into readable format.
	body, err := handler.parseResponse(response)
	if err != nil {
		return nil, err
	}

	// Return JSON marshalled response.
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.NewError(errors.JSON_DESERIALIZATION_ERROR, fmt.Sprintf("Body: %s, Error: %s", body, err.Error()))
	}

	return data, nil
}

// Create HTTP request structure.
func (handler *Handler) createHttpRequest(request *Request) (*http.Request, error) {
	var body io.Reader = nil

	// JSON marshal request body.
	if request.body != nil {
		jsonData, err := json.Marshal(request.body)

		if err != nil {
			return nil, errors.NewError(errors.JSON_SERIALIZATION_ERROR, fmt.Sprintf("Request Body: %+v, Error: %s", request.body, err.Error()))
		}

		body = bytes.NewBuffer(jsonData)
	}

	// Create HTTP request.
	httpRequest, err := http.NewRequest(string(request.method), string(request.endpoint), body)

	if err != nil {
		return nil, errors.NewError(errors.REQUEST_CREATION_ERROR, fmt.Sprintf("URL: %s, Error: %s", request.endpoint, err.Error()))
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

	return httpRequest, nil
}

// Parse response body by decompressing content according to its encoding.
// HoYoLab endpoints currently support gzip, deflate and brotli.
func (handler *Handler) parseResponse(response *http.Response) ([]byte, error) {
	var err error
	var reader io.ReadCloser

	switch encoding := response.Header.Get("Content-Encoding"); encoding {
	case string(GZIP):
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, errors.NewError(errors.RESPONSE_READ_ERROR, fmt.Sprintf("Body: %+v, Error: %s", response.Body, err.Error()))
		}

	case string(DEFLATE):
		reader = flate.NewReader(response.Body)

	default:
		reader = response.Body
	}

	body, err := io.ReadAll(reader)
	reader.Close()

	if err != nil {
		return nil, errors.NewError(errors.RESPONSE_READ_ERROR, fmt.Sprintf("Body: %+v, Error: %s", response.Body, err.Error()))
	}

	return body, nil
}
