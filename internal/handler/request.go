package handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

// HTTP request wrapper for setting parameters required for HoYoLab endpoints.
// Used by Handler to send HTTP requests.
type Request struct {
	endpoint Endpoint
	method   RequestMethod
	body     map[string]interface{}
	params   map[string]string
	headers  map[string]string
}

// Creates a new HTTP request instance.
func NewRequest(endpoint Endpoint, method RequestMethod) *Request {
	request := &Request{
		endpoint: endpoint,
		method:   method,
		body:     make(map[string]interface{}),
		params:   make(map[string]string),
		headers:  make(map[string]string),
	}

	request.setDefaultHeaders()

	return request
}

// Set body.
func (request *Request) SetBody(key string, value interface{}) {
	request.body[key] = value
}

// Set query parameter.
func (request *Request) SetParam(key string, value string) {
	request.params[key] = value
}

// Set header.
func (request *Request) SetHeader(key string, value string) {
	request.headers[key] = value
}

// Set cookie to request header.
func (request *Request) SetCookie(cookie *middleware.Cookie) {
	request.headers["Cookie"] = cookie.ParseCookie()
}

// Set language to request header.
func (request *Request) SetLanguage(language Language) {
	request.headers["X-Rpc-Language"] = string(language)
}

// Set dynamic secret to request header.
func (request *Request) SetDynamicSecret(salt DynamicSecretSalt) {
	// Generate random 6-letter string.
	length := 6
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random := make([]byte, length)

	for i := range length {
		random[i] = charset[rand.Intn(len(charset))]
	}

	// Generate MD5 hash.
	time := time.Now().Unix()
	data := []byte(fmt.Sprintf("salt=%s&t=%d&r=%s", salt, time, random))
	hash := md5.Sum(data)
	encoding := hex.EncodeToString(hash[:])

	// Parse dynamic secret and add to request header.
	ds := fmt.Sprintf("%d,%s,%s", time, random, encoding)
	request.headers["Ds"] = ds
}

// Set default headers that are required by HoYoLab endpoints.
func (request *Request) setDefaultHeaders() {
	request.headers["Accept"] = "application/json, text/plain, */*"
	request.headers["Content-Type"] = "application/json"
	request.headers["Accept-Encoding"] = fmt.Sprintf("%s, %s, %s", string(GZIP), string(DEFLATE), string(BR))
	request.headers["Sec-Ch-Ua"] = `"Chromium";v="112", "Microsoft Edge";v="112", "Not:A-Brand";v="99"`
	request.headers["Sec-Ch-Ua-Mobile"] = "?0"
	request.headers["Sec-Ch-Ua-Platform"] = `"Windows"`
	request.headers["Sec-Fetch-Dest"] = "empty"
	request.headers["Sec-Fetch-Mode"] = "cors"
	request.headers["Sec-Fetch-Site"] = "same-site"
	request.headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.46"
	request.headers["X-Rpc-App_version"] = "1.5.0"
	request.headers["X-Rpc-Client_type"] = "5"
	request.headers["X-Rpc-Language"] = string(ENGLISH)
}
