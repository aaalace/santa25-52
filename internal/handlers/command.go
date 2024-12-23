package handlers

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"santa25-52/internal/cache"
	"santa25-52/internal/context"
	"santa25-52/internal/ui"
)

func HandleCommand(ctx *context.RequestContext) {
	switch ctx.Request.Message.Command() {
	case "start":
		handleStartCommand(ctx)
	default:
		log.Println("What the fuck is this user doing with commands")
	}
}

func handleStartCommand(ctx *context.RequestContext) {
	names, _ := ctx.CacheClient.LRange(ctx.CacheClient.Context(), cache.MembersDefaultKey, 0, -1).Result()

	if len(names) == 0 {
		msg := api.NewMessage(ctx.Request.Message.Chat.ID, ui.StartPageEmpty)
		_, err := ctx.Bot.Send(msg)
		if err != nil {
			log.Println("Error sending start response [1]", err)
		}
		return
	}
	msg := api.NewMessage(ctx.Request.Message.Chat.ID, ui.StartPageMessage)
	msg.ReplyMarkup = ui.CreatePeopleListKeyboard(names)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("Error sending start response [2]", err)
	}
}
