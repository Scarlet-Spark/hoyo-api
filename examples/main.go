package main

import (
	"os"

	"github.com/Scarlet-Spark/hoyo-api/internal/middleware"
)

func main() {
	// Cookie tokens.
	ltokenV2 := os.Args[1]
	ltmidV2 := os.Args[2]
	ltuidV2 := os.Args[3]

	cookie := middleware.NewCookie(ltokenV2, ltmidV2, ltuidV2)
	GenshinDailyReward(cookie)
}
