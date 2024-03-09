package constants

// Game types in HoYoLab.
type Game string

// Daily reward endpoint parameter.
type DailyRewardParam string

const (
	// Game types.
	GAME_GENSHIN   Game = "GENSHIN"
	GAME_STAR_RAIL Game = "STAR_RAIL"

	// Daily reward endpoint parameters.
	DAILY_REWARD_HOME DailyRewardParam = "home"
	DAILY_REWARD_INFO DailyRewardParam = "info"
	DAILY_REWARD_SIGN DailyRewardParam = "sign"

	// Genshin endpoint parameters.
	GENSHIN_EVENT_ID = "sol"
	GENSHIN_ACT_ID   = "e202102251931481"

	// Star Rail endpoint parameters.
	STAR_RAIL_EVENT_ID = "luna/os"
	STAR_RAIL_ACT_ID   = "e202303301540311"

	// Languages.
	LANG_SIMPLIFIED_CHINESE = "zh-cn"
	LANG_TRADIIONAL_CHINESE = "zh-tw"
	LANG_GERMAN             = "de-de"
	LANG_ENGLISH            = "en-us"
	LANG_SPANISH            = "es-es"
	LANG_FRENCH             = "fr-fr"
	LANG_INDONESIAN         = "id-id"
	LANG_ITALIAN            = "it-it"
	LANG_JAPANESE           = "ja-jp"
	LANG_KOREAN             = "ko-kr"
	LANG_PORTUGUESE         = "pt-pt"
	LANG_RUSSIAN            = "ru-ru"
	LANG_THAI               = "th-th"
	LANG_TURKISH            = "tr-tr"
	LANG_VIETNAMESE         = "vi-vn"

	// Dynamic secret salts.
	DS_GLOBAL    = "6s25p5ox5y14umn1p61aqyyvbvvl3lrt"
	DS_APP_LOGIN = "IZPgfb0dRPtBeLuFkdDznSZ6f4wWt6y2"
)

// Some endpoints are shared across different games with only minor differences to the URL e.g., daily rewards.
// This struct consolidates the common differences between each HoYoverse game.
type gameParams struct {
	baseUrl string
	eventId string
	actId   string
}

// Specific endpoint parameters for each game.
var gameEndpointParams = map[Game]gameParams{
	GAME_GENSHIN: {
		baseUrl: HK4E_API,
		eventId: GENSHIN_EVENT_ID,
		actId:   GENSHIN_ACT_ID,
	},
	GAME_STAR_RAIL: {
		baseUrl: PUBLIC_API,
		eventId: STAR_RAIL_EVENT_ID,
		actId:   STAR_RAIL_ACT_ID,
	},
}
