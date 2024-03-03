package handler

// Language codes for localization on HoYoLab endpoints.
type Language string

// Salts for generating dynamic secrets.
type DynamicSecretSalt string

// Request methods.
type RequestMethod string

// Content encoding formats.
type ContentEncoding string

const (
	// Languages.
	SIMPLIFIED_CHINESE Language = "zh-cn"
	TRADIIONAL_CHINESE Language = "zh-tw"
	GERMAN             Language = "de-de"
	ENGLISH            Language = "en-us"
	SPANISH            Language = "es-es"
	FRENCH             Language = "fr-fr"
	INDONESIAN         Language = "id-id"
	ITALIAN            Language = "it-it"
	JAPANESE           Language = "ja-jp"
	KOREAN             Language = "ko-kr"
	PORTUGUESE         Language = "pt-pt"
	RUSSIAN            Language = "ru-ru"
	THAI               Language = "th-th"
	TURKISH            Language = "tr-tr"
	VIETNAMESE         Language = "vi-vn"

	// Dynamic secret salts.
	GLOBAL    DynamicSecretSalt = "6s25p5ox5y14umn1p61aqyyvbvvl3lrt"
	APP_LOGIN DynamicSecretSalt = "IZPgfb0dRPtBeLuFkdDznSZ6f4wWt6y2"

	// Request methods.
	GET  RequestMethod = "GET"
	POST RequestMethod = "POST"

	// Content encodings.
	GZIP    ContentEncoding = "gzip"
	DEFLATE ContentEncoding = "deflate"
	BR      ContentEncoding = "br"
)

// Language stringer.
func (language Language) String() string {
	return string(language)
}

// Dynamic secret salt stringer.
func (salt DynamicSecretSalt) String() string {
	return string(salt)
}

// Request method stringer.
func (method RequestMethod) String() string {
	return string(method)
}

// Content encoding stringer.
func (encoding ContentEncoding) String() string {
	return string(encoding)
}
