package main

import (
	"github.com/Scarlet-Spark/hoyo-api/client"
	"github.com/Scarlet-Spark/hoyo-api/internal/constants"
	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

// Call daily reward API using Genshin client.
func GenshinDailyReward(cookie middleware.Cookie) {
	// Genshin daily rewards.
	genshin := client.NewGenshinClient(cookie, constants.LANG_ENGLISH, 800874180)
	response, err := genshin.Daily.Claim()
	printResponse(response, err, false)
}
