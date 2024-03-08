package middleware

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	// Token keys.
	ltokenV2Key = "ltoken_v2"
	ltmidV2Key  = "ltmid_v2"
	ltuidV2Key  = "ltuid_v2"
)

// Cookie data class that stores tokens required for calling HoYoLab endpoints.
type Cookie struct {
	tokens map[string]http.Cookie
}

// Constructor.
func NewCookie(ltokenV2 string, ltmidV2 string, ltuidV2 string) Cookie {
	ltokenV2Cookie := http.Cookie{Name: ltokenV2Key, Value: ltokenV2}
	ltmidV2Cookie := http.Cookie{Name: ltmidV2Key, Value: ltmidV2}
	ltuidV2Cookie := http.Cookie{Name: ltuidV2Key, Value: ltuidV2}

	tokens := map[string]http.Cookie{
		ltokenV2Key: ltokenV2Cookie,
		ltmidV2Key:  ltmidV2Cookie,
		ltuidV2Key:  ltuidV2Cookie,
	}

	return Cookie{tokens: tokens}
}

// Converts cookie tokens into a string to be added into a request header.
// Each token is separated by semicolons.
// Similar to http.Request.AddCookie().
func (cookie Cookie) ParseCookie() string {
	count := 0
	length := len(cookie.tokens)
	buffer := new(bytes.Buffer)

	for key, value := range cookie.tokens {
		count++
		fmt.Fprintf(buffer, "%s=%s", key, value.Value)
		if count != length {
			fmt.Fprint(buffer, ";")
		}
	}

	return buffer.String()
}

// Get all cookies.
func (cookie Cookie) GetCookies() []http.Cookie {
	tokens := make([]http.Cookie, len(cookie.tokens))

	for _, token := range cookie.tokens {
		tokens = append(tokens, token)
	}

	return tokens
}
