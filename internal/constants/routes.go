package constants

import "fmt"

const (
	// Main API endpoints.
	BBS_API    = "https://bbs-api-os.hoyolab.com"
	HK4E_API   = "https://sg-hk4e-api.hoyolab.com"
	PUBLIC_API = "https://sg-public-api.hoyolab.com"
	WIKI_API   = "https://sg-wiki-api-static.hoyolab.com"

	// Genshin Impact Battle Chronicle API endpoints.
	GENSHIN_RECORD_INDEX_API        = BBS_API + "/game_record/genshin/api/index"
	GENSHIN_RECORD_SPIRAL_ABYSS_API = BBS_API + "/game_record/genshin/api/spiralAbyss"

	// HoYoWiki API endpoints.
	HOYOWIKI_ENTRY_PAGE_LIST_API = WIKI_API + "/hoyowiki/wapi/get_entry_page_list"
	HOYOWIKI_ENTRY_PAGE_API      = WIKI_API + "/hoyowiki/wapi/entry_page"
)

// Returns the endpoint for daily rewards.
func DailyRewardAPI(game Game, daily DailyRewardParam) string {
	params := gameEndpointParams[game]
	return fmt.Sprintf("%s/event/%s/%s?act_id=%s", params.baseUrl, params.eventId, string(daily), params.actId)
}
