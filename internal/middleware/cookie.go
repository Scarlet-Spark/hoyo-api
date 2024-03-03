package middleware

import (
	"bytes"
	"fmt"
)

const (
	// Token keys.
	LTOKEN_V2_KEY = "ltoken_v2"
	LTMID_V2_KEY  = "ltmid_v2"
	LTUID_V2_KEY  = "ltuid_v2"
)

// Cookie data class that stores tokens required for calling HoYoLab endpoints.
type Cookie struct {
	tokens map[string]string
}

// Creates a new cookie instance.
func NewCookie(ltokenV2 string, ltmidV2 string, ltuidV2 string) *Cookie {
	tokens := map[string]string{
		LTOKEN_V2_KEY: ltokenV2,
		LTMID_V2_KEY:  ltmidV2,
		LTUID_V2_KEY:  ltuidV2,
	}

	cookie := &Cookie{tokens}
	return cookie
}

// Converts cookie tokens into a string used when sending requests.
func (cookie *Cookie) ParseCookie() string {
	count := 0
	length := len(cookie.tokens)
	buffer := new(bytes.Buffer)

	for key, value := range cookie.tokens {
		count++
		fmt.Fprintf(buffer, "%s=%s", key, value)
		if count != length {
			fmt.Fprint(buffer, ";")
		}
	}

	return buffer.String()
}
