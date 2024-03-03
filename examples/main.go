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

	request := handler.NewRequest(handler.GENSHIN_RECORD_SPIRAL_ABYSS_API, handler.GET)
	request.SetCookie(cookie)
	request.SetParam("role_id", "800874180")
	request.SetParam("server", "os_asia")
	request.SetParam("schedule_type", "1")
	request.SetDynamicSecret(handler.GLOBAL)

	requestHandler := handler.NewHandler()
	response, err := requestHandler.Send(request)

	if err != nil {
		fmt.Println(err)
	}

	data, _ := json.MarshalIndent(response, "", "    ")
	fmt.Println(string(data))
}
