package main

import (
	bp "santa25-52/internal/bot"
	"santa25-52/internal/config"
	"santa25-52/internal/context"
	"santa25-52/internal/db"
)

func main() {
	cfg := config.MustLoad()
	bot := bp.MustLoad(cfg)
	dbClient := db.MustLoad(cfg)
	ctx := context.BuildBaseContext(bot.Bot, dbClient)

	// permanently recovers
	bot.SafeLaunch(ctx)
}
