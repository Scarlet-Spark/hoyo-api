package handler

// HoYoLab API endpoints.
type Endpoint string

const (
	// Main API endpoints.
	BBS_API  Endpoint = "https://bbs-api-os.hoyolab.com"
	HK4E_API Endpoint = "https://sg-hk4e-api.hoyolab.com"
	WIKI_API Endpoint = "https://sg-wiki-api-static.hoyolab.com"

	// Genshin Impact Battle Chronicle API endpoints.
	GENSHIN_RECORD_INDEX_API        Endpoint = BBS_API + "/game_record/genshin/api/index"
	GENSHIN_RECORD_SPIRAL_ABYSS_API Endpoint = BBS_API + "/game_record/genshin/api/spiralAbyss"

	// Genshin Impact Daily Check-In API endpoints.
	GENSHIN_DAILY_CHECK_IN_EVENT_ID  Endpoint = "sol"
	GENSHIN_DAILY_CHECK_IN_ACT_ID    Endpoint = "e202102251931481"
	GENSHIN_DAILY_CHECK_IN_HOME_API  Endpoint = HK4E_API + "/event/" + GENSHIN_DAILY_CHECK_IN_EVENT_ID + "/home?act_id=" + GENSHIN_DAILY_CHECK_IN_ACT_ID
	GENSHIN_DAILY_CHECK_IN_INFO_API  Endpoint = HK4E_API + "/event/" + GENSHIN_DAILY_CHECK_IN_EVENT_ID + "/info?act_id=" + GENSHIN_DAILY_CHECK_IN_ACT_ID
	GENSHIN_DAILY_CHECK_IN_CLAIM_API Endpoint = HK4E_API + "/event/" + GENSHIN_DAILY_CHECK_IN_EVENT_ID + "/sign?act_id=" + GENSHIN_DAILY_CHECK_IN_ACT_ID

	// HoYoWiki API endpoints.
	HOYOWIKI_ENTRY_LIST_API Endpoint = WIKI_API + "/hoyowiki/wapi/get_entry_page_list"
	HOYOWIKI_ENTRY_API      Endpoint = WIKI_API + "hoyowiki/wapi/entry_page"
)
