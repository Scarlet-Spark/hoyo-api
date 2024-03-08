package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Scarlet-Spark/hoyo-api/internal/constants"
	"github.com/Scarlet-Spark/hoyo-api/internal/handler"
	"github.com/Scarlet-Spark/hoyo-api/internal/interfaces"
	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

func main() {
	// Cookie tokens.
	ltokenV2 := ""
	ltmidV2 := ""
	ltuidV2 := ""

	cookie := middleware.NewCookie(ltokenV2, ltmidV2, ltuidV2)

	// // // Get Spiral Abyss stats.
	request := handler.NewRequest(constants.GENSHIN_RECORD_SPIRAL_ABYSS_API, http.MethodGet).
		AddCookie(cookie).
		AddParam("role_id", "800874180").
		AddParam("server", "os_asia").
		AddParam("schedule_type", "1").
		AddDynamicSecret(constants.DS_GLOBAL).
		AddLanguage(constants.LANG_ENGLISH).
		Build()

	// Get Genshin characters list in HoYoWiki.
	// request := handler.NewRequest(constants.HOYOWIKI_ENTRY_PAGE_LIST_API, http.MethodPost).
	// 	AddReferer("https://wiki.hoyolab.com").
	// 	AddBody("filters", []string{}).
	// 	AddBody("menu_id", 2).    // Genshin Character List
	// 	AddBody("page_num", 1).   // Pagination
	// 	AddBody("page_size", 30). // Number of items returned
	// 	AddBody("use_es", true).
	// 	Build()

	requestHandler := handler.NewHandler(cookie)

	response, err := requestHandler.Send(request)
	printResult(response, err, true)

	// Daily rewards.
	daily := interfaces.NewDailyRewardComponent(
		constants.HK4E_API,
		constants.GENSHIN_DAILY_CHECK_IN_EVENT_ID,
		constants.GENSHIN_DAILY_CHECK_IN_ACT_ID,
		constants.LANG_ENGLISH,
		requestHandler,
	)

	response, err = daily.Claim()
	printResult(response, err, false)

}

// Print result.
func printResult(response map[string]interface{}, err error, ignore bool) {
	if err != nil {
		fmt.Println(err)
	} else {
		if ignore {
			return
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	}
}
