package main

import (
	bp "santa25-52/internal/bot"
	"santa25-52/internal/cache"
	"santa25-52/internal/config"
	"santa25-52/internal/context"
)

func main() {
	cfg := config.MustLoad()

	bot := bp.MustLoad(cfg)
	cacheClient := cache.MustLoad(cfg)

	ctx := context.BuildBaseContext(bot.Bot, cacheClient)

	// permanently recovers
	bot.SafeLaunch(ctx)
}
