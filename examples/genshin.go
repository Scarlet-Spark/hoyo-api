package main

import (
	"github.com/Scarlet-Spark/hoyo-api/client"
	"github.com/Scarlet-Spark/hoyo-api/internal/constants"
	"github.com/Scarlet-Spark/hoyo-api/internal/handler"
	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

// Call daily reward API using Genshin client.
func GenshinDailyReward(cookie middleware.Cookie) {
	handler := handler.NewHandler(cookie)

	// Genshin daily rewards.
	genshin := client.NewGenshinClient(handler, constants.LANG_ENGLISH)
	response, err := genshin.Daily.Claim()
	printResult(response, err, false)
}
