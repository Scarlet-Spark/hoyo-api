package interfaces

import (
	"net/http"

	"github.com/Scarlet-Spark/hoyo-api/internal/constants"
	"github.com/Scarlet-Spark/hoyo-api/internal/handler"
)

// Daily reward component.
// Responsible for getting information of current daily reward status and claiming daily reward.
type DailyReward struct {
	game     constants.Game
	language string
	handler  handler.Handler
}

// Constructor.
func NewDailyReward(game constants.Game, language string, handler handler.Handler) DailyReward {
	return DailyReward{
		game:     game,
		language: language,
		handler:  handler,
	}
}

// Get the list of available daily rewards for the month.
func (daily DailyReward) List() (map[string]interface{}, error) {
	endpoint := constants.DailyRewardAPI(daily.game, constants.DAILY_REWARD_HOME)
	request := handler.NewRequest(endpoint, http.MethodGet).
		AddLanguage(daily.language).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.handler.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Get information on the current daily reward status.
func (daily DailyReward) Info() (map[string]interface{}, error) {
	endpoint := constants.DailyRewardAPI(daily.game, constants.DAILY_REWARD_INFO)
	request := handler.NewRequest(endpoint, http.MethodGet).
		AddLanguage(daily.language).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.handler.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Claim daily reward.
func (daily DailyReward) Claim() (map[string]interface{}, error) {
	endpoint := constants.DailyRewardAPI(daily.game, constants.DAILY_REWARD_SIGN)
	request := handler.NewRequest(endpoint, http.MethodPost).
		AddLanguage(daily.language).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.handler.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}
