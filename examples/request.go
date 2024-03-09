package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Scarlet-Spark/hoyo-api/internal/constants"
	"github.com/Scarlet-Spark/hoyo-api/internal/handler"
	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

// Get Spiral Abyss information.
func GetSpiralAbyssInfo(cookie middleware.Cookie) {
	request := handler.NewRequest(constants.GENSHIN_RECORD_SPIRAL_ABYSS_API, http.MethodGet).
		AddCookie(cookie).
		AddParam("role_id", "800874180").
		AddParam("server", "os_asia").
		AddParam("schedule_type", "1").
		AddDynamicSecret(constants.DS_GLOBAL).
		AddLanguage(constants.LANG_ENGLISH).
		Build()

	handler := handler.NewHandler(cookie)

	response, err := handler.Send(request)
	printResult(response, err, false)
}

// Get Genshin characters list in HoYoWiki.
func GetGenshinCharacters(cookie middleware.Cookie) {
	request := handler.NewRequest(constants.HOYOWIKI_ENTRY_PAGE_LIST_API, http.MethodPost).
		AddReferer("https://wiki.hoyolab.com").
		AddBody("filters", []string{}).
		AddBody("menu_id", 2).    // Genshin Character List
		AddBody("page_num", 1).   // Pagination
		AddBody("page_size", 30). // Number of items returned
		AddBody("use_es", true).
		Build()

	handler := handler.NewHandler(cookie)

	response, err := handler.Send(request)
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
