package interfaces

import (
	"net/http"

	"github.com/Scarlet-Spark/hoyo-api/internal/constants"
	"github.com/Scarlet-Spark/hoyo-api/internal/handler"
)

// Daily reward component.
// Responsible for getting information of current daily reward status and claiming daily reward.
type DailyRewardComponent struct {
	baseUrl  string
	eventId  string
	actId    string
	language string
	handler.Handler
}

// Constructor.
func NewDailyRewardComponent(baseUrl string, eventId string, actId string, language string, handler handler.Handler) DailyRewardComponent {
	return DailyRewardComponent{
		baseUrl:  baseUrl,
		eventId:  eventId,
		actId:    actId,
		language: language,
		Handler:  handler,
	}
}

// Get the list of available daily rewards for the month.
func (daily DailyRewardComponent) List() (map[string]interface{}, error) {
	endpoint := constants.DailyRewardListAPI(daily.baseUrl, daily.eventId, daily.actId)
	request := handler.NewRequest(endpoint, http.MethodGet).
		AddLanguage(daily.language).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Get information on the current daily reward status.
func (daily DailyRewardComponent) Info() (map[string]interface{}, error) {
	endpoint := constants.DailyRewardInfoAPI(daily.baseUrl, daily.eventId, daily.actId)
	request := handler.NewRequest(endpoint, http.MethodGet).
		AddLanguage(daily.language).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Claim daily reward.
func (daily DailyRewardComponent) Claim() (map[string]interface{}, error) {
	endpoint := constants.DailyRewardClaimAPI(daily.baseUrl, daily.eventId, daily.actId)
	request := handler.NewRequest(endpoint, http.MethodPost).
		AddLanguage(daily.language).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}
