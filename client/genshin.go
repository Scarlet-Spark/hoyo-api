package client

import (
	"fmt"
	"net/http"

	"github.com/Scarlet-Spark/hoyo-api/internal/constants"
	"github.com/Scarlet-Spark/hoyo-api/internal/errors"
	"github.com/Scarlet-Spark/hoyo-api/internal/handler"
	"github.com/Scarlet-Spark/hoyo-api/internal/interfaces"
	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

// Client that interfaces to HoYoLab endpoints related to Genshin Impact.
// i.e., Spiral Abyss, Daily Reward
type GenshinClient struct {
	cache    *middleware.Cache
	handler  handler.Handler
	Language constants.Language
	UserId   int
	Daily    interfaces.DailyReward
}

// Constructor.
func NewGenshinClient(cookie middleware.Cookie, language constants.Language, uid int) *GenshinClient {
	handler := handler.NewHandler(cookie)

	return &GenshinClient{
		cache:    middleware.NewCache(),
		handler:  handler,
		Language: language,
		UserId:   uid,
		Daily: interfaces.NewDailyReward(
			constants.GAME_GENSHIN,
			language,
			handler,
		),
	}
}

// Get current Spiral Abyss information.
// Set current argument to false to get previous cycle's information.
func (genshin GenshinClient) SpiralAbyss(current bool) (map[string]interface{}, error) {
	scheduleType := 1
	if !current {
		scheduleType = 2
	}

	server, err := genshin.getRegion()
	if err != nil {
		return nil, err
	}

	request := handler.NewRequest(constants.GENSHIN_RECORD_SPIRAL_ABYSS_API, http.MethodGet).
		AddParam("role_id", fmt.Sprint(genshin.UserId)).
		AddParam("server", string(*server)).
		AddParam("schedule_type", fmt.Sprint(scheduleType)).
		AddDynamicSecret(constants.DS_GLOBAL).
		AddLanguage(constants.LANG_ENGLISH).
		Build()

	data, err := genshin.handler.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Get Genshin regional server code based off the user ID.
func (genshin GenshinClient) getRegion() (*constants.GenshinRegion, error) {
	// Get 1st digit of user ID to determine the region.
	i := genshin.UserId
	for i >= 10 {
		i /= 10
	}

	regions := map[int]constants.GenshinRegion{
		6: constants.GENSHIN_REGION_USA,
		7: constants.GENSHIN_REGION_EUROPE,
		8: constants.GENSHIN_REGION_ASIA,
		9: constants.GENSHIN_REGION_CHINA_TAIWAN,
	}

	region, exists := regions[i]
	if !exists {
		return nil,
			errors.NewError(
				errors.REGION_SERVER_CODE_ERROR,
				fmt.Sprintf("Genshin regional server code not found for UID %d", genshin.UserId),
			)
	}

	return &region, nil
}
