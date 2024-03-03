package handler

// Language codes for localization on HoYoLab endpoints.
type LanguageEnum string

// Salts for generating dynamic secrets.
type DsSaltEnum string

// Request methods.
type RequestMethod string

// Content encoding formats.
type ContentEncoding string

const (
	// Languages.
	SIMPLIFIED_CHINESE LanguageEnum = "zh-cn"
	TRADIIONAL_CHINESE LanguageEnum = "zh-tw"
	GERMAN             LanguageEnum = "de-de"
	ENGLISH            LanguageEnum = "en-us"
	SPANISH            LanguageEnum = "es-es"
	FRENCH             LanguageEnum = "fr-fr"
	INDONESIAN         LanguageEnum = "id-id"
	ITALIAN            LanguageEnum = "it-it"
	JAPANESE           LanguageEnum = "ja-jp"
	KOREAN             LanguageEnum = "ko-kr"
	PORTUGUESE         LanguageEnum = "pt-pt"
	RUSSIAN            LanguageEnum = "ru-ru"
	THAI               LanguageEnum = "th-th"
	TURKISH            LanguageEnum = "tr-tr"
	VIETNAMESE         LanguageEnum = "vi-vn"

	// Dynamic secret salts.
	GLOBAL    DsSaltEnum = "6s25p5ox5y14umn1p61aqyyvbvvl3lrt"
	APP_LOGIN DsSaltEnum = "IZPgfb0dRPtBeLuFkdDznSZ6f4wWt6y2"

	// Request methods.
	GET  RequestMethod = "GET"
	POST RequestMethod = "POST"

	// Content encodings.
	GZIP    ContentEncoding = "gzip"
	DEFLATE ContentEncoding = "deflate"
	BR      ContentEncoding = "br"
)
