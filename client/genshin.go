package client

import (
	"github.com/Scarlet-Spark/hoyo-api/internal/constants"
	"github.com/Scarlet-Spark/hoyo-api/internal/handler"
	"github.com/Scarlet-Spark/hoyo-api/internal/interfaces"
	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

// Client that interfaces to HoYoLab endpoints related to Genshin Impact.
// i.e., Spiral Abyss, Daily Reward
type GenshinClient struct {
	cache    middleware.Cache
	handler  handler.Handler
	Language string
	Daily    interfaces.DailyReward
}

// Constructor.
func NewGenshinClient(handler handler.Handler, language string) *GenshinClient {
	return &GenshinClient{
		cache:    *middleware.NewCache(),
		handler:  handler,
		Language: language,
		Daily: interfaces.NewDailyReward(
			constants.GAME_GENSHIN,
			constants.LANG_ENGLISH,
			handler,
		),
	}
}
