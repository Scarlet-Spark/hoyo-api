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

	// Daily Rewards parameters.
	GENSHIN_DAILY_CHECK_IN_EVENT_ID = "sol"
	GENSHIN_DAILY_CHECK_IN_ACT_ID   = "e202102251931481"
)

// Returns the endpoint for daily reward list.
func DailyRewardListAPI(baseUrl string, eventId string, actId string) string {
	return fmt.Sprintf("%s/event/%s/home?act_id=%s", baseUrl, eventId, actId)
}

// Returns the endpoint for daily reward info.
func DailyRewardInfoAPI(baseUrl string, eventId string, actId string) string {
	return fmt.Sprintf("%s/event/%s/info?act_id=%s", baseUrl, eventId, actId)
}

// Returns the endpoint for daily reward claim.
func DailyRewardClaimAPI(baseUrl string, eventId string, actId string) string {
	return fmt.Sprintf("%s/event/%s/sign?act_id=%s", baseUrl, eventId, actId)
}
