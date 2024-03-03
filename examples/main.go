package main

import (
	"encoding/json"
	"fmt"

	"github.com/Scarlet-Spark/hoyo-api/internal/handler"
	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

func main() {
	// Cookie tokens.
	ltokenV2 := ""
	ltmidV2 := ""
	ltuidV2 := ""

	cookie := middleware.NewCookie(ltokenV2, ltmidV2, ltuidV2)

	// Get Spiral Abyss stats.
	request := handler.NewRequest(handler.GENSHIN_RECORD_SPIRAL_ABYSS_API, handler.GET)
	request.SetCookie(cookie)
	request.SetParam("role_id", "800874180")
	request.SetParam("server", "os_asia")
	request.SetParam("schedule_type", "1")
	request.SetDynamicSecret(handler.GLOBAL)

	// Get Genshin characters list in HoYoWiki.
	request = handler.NewRequest(handler.HOYOWIKI_ENTRY_PAGE_LIST_API, handler.POST)
	request.SetHeader("Referer", "https://wiki.hoyolab.com")
	request.SetBody("filters", []string{})
	request.SetBody("menu_id", 2)    // Genshin Character List
	request.SetBody("page_num", 1)   // Pagination
	request.SetBody("page_size", 30) // Number of items returned
	request.SetBody("use_es", true)

	ignore := false
	requestHandler := handler.NewHandler()

	response, err := requestHandler.Send(request)

	if err.Exists {
		fmt.Println(err)
	} else {
		if ignore {
			return
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	}
}
