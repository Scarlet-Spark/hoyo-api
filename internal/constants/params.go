package constants

// Game types in HoYoLab.
type Game string

// Daily reward endpoint parameter.
type DailyRewardParam string

// Language codes supported in HoYoLab.
type Language string

// Genshin regional server codes.
type GenshinRegion string

// Dynamic secret salts used in HoYoLab.
type DynamicSecret string

const (
	// Game types.
	GAME_GENSHIN   Game = "GENSHIN"
	GAME_STAR_RAIL Game = "STAR_RAIL"

	// Daily reward endpoint parameters.
	DAILY_REWARD_HOME DailyRewardParam = "home"
	DAILY_REWARD_INFO DailyRewardParam = "info"
	DAILY_REWARD_SIGN DailyRewardParam = "sign"

	// Languages.
	LANG_SIMPLIFIED_CHINESE Language = "zh-cn"
	LANG_TRADIIONAL_CHINESE Language = "zh-tw"
	LANG_GERMAN             Language = "de-de"
	LANG_ENGLISH            Language = "en-us"
	LANG_SPANISH            Language = "es-es"
	LANG_FRENCH             Language = "fr-fr"
	LANG_INDONESIAN         Language = "id-id"
	LANG_ITALIAN            Language = "it-it"
	LANG_JAPANESE           Language = "ja-jp"
	LANG_KOREAN             Language = "ko-kr"
	LANG_PORTUGUESE         Language = "pt-pt"
	LANG_RUSSIAN            Language = "ru-ru"
	LANG_THAI               Language = "th-th"
	LANG_TURKISH            Language = "tr-tr"
	LANG_VIETNAMESE         Language = "vi-vn"

	// Genshin regional server codes.
	GENSHIN_REGION_USA          GenshinRegion = "os_usa"
	GENSHIN_REGION_EUROPE       GenshinRegion = "os_euro"
	GENSHIN_REGION_ASIA         GenshinRegion = "os_asia"
	GENSHIN_REGION_CHINA_TAIWAN GenshinRegion = "os_cht"

	// Dynamic secret salts.
	DS_GLOBAL    DynamicSecret = "6s25p5ox5y14umn1p61aqyyvbvvl3lrt"
	DS_APP_LOGIN DynamicSecret = "IZPgfb0dRPtBeLuFkdDznSZ6f4wWt6y2"
)

const (
	// Genshin endpoint parameters.
	genshinEventId = "sol"
	genshinActId   = "e202102251931481"

	// Star Rail endpoint parameters.
	starRailEventId = "luna/os"
	starRailActId   = "e202303301540311"
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
		eventId: genshinEventId,
		actId:   genshinActId,
	},
	GAME_STAR_RAIL: {
		baseUrl: PUBLIC_API,
		eventId: starRailEventId,
		actId:   starRailActId,
	},
}
